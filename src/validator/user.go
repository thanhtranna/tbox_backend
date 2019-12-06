package validator

import (
	"fmt"
	"regexp"
)

type IUserValidate interface {
	ValidatePhoneNumber(string) error
	ValidateUserID(string) error
}

const (
	regexPhoneNumber string = `^(09|01[2|6|8|9])+([0-9]{8})$`
	regexUUID        string = `^[0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12}$`
)

type userValidator struct{}

func NewUserValidator() IUserValidate {
	return &userValidator{}
}

func (u *userValidator) ValidatePhoneNumber(input string) error {
	regex := regexp.MustCompile(regexPhoneNumber)
	if !regex.MatchString(input) {
		return fmt.Errorf("phone number invalid")
	}
	return nil
}

func (u *userValidator) ValidateUserID(userID string) error {
	regex := regexp.MustCompile(regexUUID)
	if !regex.MatchString(userID) {
		return fmt.Errorf("user id invalid")
	}

	return nil
}
