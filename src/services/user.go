package services

import (
	"fmt"
	"tbox_backend/src/entity"

	"tbox_backend/src/common"
	"tbox_backend/src/helper"
	"tbox_backend/src/repository"
	"tbox_backend/src/validator"
)

//IUserService: interface user service
type IUserService interface {
	Login(string) (string, error)
	VerifyPhoneNumber(entity.VerifyPhoneNumber) (string, error)
	ResendOTP(entity.ResendOTP) error
}

type userService struct {
	userRepo      repository.IUserRepo
	userTokenRepo repository.IUserTokenRepo
	userOTPRepo   repository.IUserOTPRepo
	twilioHelper  helper.ITwilioHelper
	authHelper    helper.IAuthHelper
	userValidator validator.IUserValidate
}

func NewUserService(userRepo repository.IUserRepo, userTokenRepo repository.IUserTokenRepo, userOTPRepo repository.IUserOTPRepo, twilioHelper helper.ITwilioHelper, authHelper helper.IAuthHelper, userValidator validator.IUserValidate) IUserService {
	return &userService{
		userRepo:      userRepo,
		userTokenRepo: userTokenRepo,
		userOTPRepo:   userOTPRepo,
		twilioHelper:  twilioHelper,
		authHelper:    authHelper,
		userValidator: userValidator,
	}
}

func (u *userService) Login(input string) (string, error) {
	if err := u.userValidator.ValidatePhoneNumber(input); err != nil {
		return "", err
	}

	// add prefix phone number
	phoneNumber := common.AddPrefixPhoneNumberVietNam(input)

	// find user by phone number
	user, err := u.userRepo.FindByPhoneNumber(phoneNumber)
	if err != nil {
		return "", err
	}

	// verify phone number
	if user.IsVerify == true {
		userToken, err := u.userTokenRepo.GetTokenByUserID(user.ID)
		if err != nil {
			return "", err
		}
		// return token and end
		return userToken.Token, nil
	}
	// Check flag block user sended otp
	if _, ok := u.userOTPRepo.CheckSendedOTP(user.ID); ok {
		return "", fmt.Errorf("Please enter OTP")
	}

	// generation OTP with phone number
	otp, err := common.GetTOTPToken(common.String(16))
	if err != nil {
		return "", err
	}

	// Send OTP to user
	err = u.twilioHelper.SendOTP(otp, user.PhoneNumber)
	if err != nil {
		return "", err
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
		return "", err
	}

	return user.ID, nil
}

func (u *userService) VerifyPhoneNumber(input entity.VerifyPhoneNumber) (string, error) {
	if err := u.userValidator.ValidateVerifyPhoneNumber(input); err != nil {
		return "", err
	}

	// find user by userId
	user, err := u.userRepo.FindByUserID(input.UserID)
	if err != nil {
		return "", err
	}
	if user.IsVerify == true {
		return "", fmt.Errorf("phone number is verified")
	}
	otp, ok := u.userOTPRepo.CheckSendedOTP(input.UserID)
	if !ok {
		return "", fmt.Errorf("OTP expire, Please resend OTP")
	}

	if otp != input.OTP {
		return "", fmt.Errorf("OTP not match. Please enter OTP again")
	}
	// OTP correct
	user.IsVerify = true
	user.UpdatedAt = common.GetVietNamTime()

	err = u.userRepo.UpdateVerifyPhoneNumber(user)
	if err != nil {
		return "", err
	}
	token, err := u.authHelper.GenerateToken(user.PhoneNumber, user.ID)
	if err != nil {
		return "", err
	}
	err = u.userTokenRepo.Save(entity.UserToken{
		ID:        user.ID,
		Token:     token,
		CreatedAt: common.GetVietNamTime(),
		UpdatedAt: common.GetVietNamTime(),
	})
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *userService) ResendOTP(input entity.ResendOTP) error {
	if err := u.userValidator.ValidateUserID(input.UserID); err != nil {
		return err
	}
	user, err := u.userRepo.FindByUserID(input.UserID)
	if err != nil {
		return err
	}
	if user.IsVerify == true {
		return fmt.Errorf("phone number has verified")
	}
	// generation OTP with phone number
	otp, err := common.GetTOTPToken(common.String(16))
	if err != nil {
		return err
	}
	// Send OTP to user
	err = u.twilioHelper.SendOTP(otp, user.PhoneNumber)
	if err != nil {
		return err
	}
	_ = u.userOTPRepo.CacheOTPWithUser(user.ID, otp)
	return nil
}
