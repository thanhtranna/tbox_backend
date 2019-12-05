package services

import (
	"fmt"
	"tbox_backend/src/repository"
	"tbox_backend/src/validator"
)

type IUserService interface {
	Login(string) error
}

type userService struct {
	userRepo      repository.IUserRepo
	userTokenRepo repository.IUserTokenRepo
}

func NewUserService(userRepo repository.IUserRepo, userTokenRepo repository.IUserTokenRepo) IUserService {
	return &userService{
		userRepo:      userRepo,
		userTokenRepo: userTokenRepo,
	}
}

func (u *userService) Login(input string) error {
	var userValidator = validator.NewUserValidator()
	if err := userValidator.ValidatePhoneNumber(input); err != nil {
		return err
	}

	data, err := u.userRepo.FindByPhoneNumber(input)
	if err != nil {
		return err
	}

	// verify phone number
	if data.IsVerify == true {
		token, err := u.userTokenRepo.GetTokenByUserID(data.ID)
		if err != nil {
			return err
		}
		// return token if login
		fmt.Println("token", token)
	}

	fmt.Println("data ", data)

	return nil
}
