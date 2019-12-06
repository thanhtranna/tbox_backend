package services

import (
	"fmt"
	"tbox_backend/src/entity"

	"tbox_backend/src/common"
	"tbox_backend/src/helper"
	"tbox_backend/src/repository"
	"tbox_backend/src/validator"
)

type IUserService interface {
	Login(string) error
}

type userService struct {
	userRepo      repository.IUserRepo
	userTokenRepo repository.IUserTokenRepo
	userOTPRepo   repository.IUserOTPRepo
	twilioHelper  helper.ITwilioHelper
}

func NewUserService(userRepo repository.IUserRepo, userTokenRepo repository.IUserTokenRepo, userOTPRepo repository.IUserOTPRepo, twilioHelper helper.ITwilioHelper) IUserService {
	return &userService{
		userRepo:      userRepo,
		userTokenRepo: userTokenRepo,
		userOTPRepo:   userOTPRepo,
		twilioHelper:  twilioHelper,
	}
}

func (u *userService) Login(input string) error {
	var userValidator = validator.NewUserValidator()
	if err := userValidator.ValidatePhoneNumber(input); err != nil {
		return err
	}

	// add prefix phone number
	phoneNumber := common.AddPrefixPhoneNumberVietNam(input)

	user, err := u.userRepo.FindByPhoneNumber(phoneNumber)
	if err != nil {
		fmt.Println("err1", err)
		return err
	}

	// verify phone number
	if user.IsVerify == true {
		token, err := u.userTokenRepo.GetTokenByUserID(user.ID)
		if err != nil {
			fmt.Println("err2", err)
			return err
		}
		// return token if login
		// return token and end
		fmt.Println("token", token)
		return nil
	}
	// Todo: check xem da gui ma OTP chua?
	// Check flag block user sended otp
	if _, ok := u.userOTPRepo.CheckSendedOTP(user.ID); ok {
		fmt.Println("Dang cho xac thuc phone number")
		return nil
	}

	// generation OTP with phone number
	otp, err := common.GetTOTPToken(common.String(16))
	if err != nil {
		return err
	}
	// chua gui hoac la chua dang nhap lan nao
	// TODO: gui tin nhan den cho user trong vong 60s.
	err = u.twilioHelper.SendOTP(otp, user.PhoneNumber)
	if err != nil {
		return err
	}

	// Cache and insert into mongo
	_ = u.userOTPRepo.CacheOTPWithUser(user.ID, otp)
	err = u.userOTPRepo.Save(entity.UserOTP{
		ID:         user.ID,
		OTP:        otp,
		TimeExpire: common.GetVietNamTime() + int64(common.TimeCacheOTP),
		CreatedAt:  common.GetVietNamTime(),
		UpdatedAt:  common.GetVietNamTime(),
	})
	if err != nil {
		return err
	}
	fmt.Println("user ", user)

	return nil
}
