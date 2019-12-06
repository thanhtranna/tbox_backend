package validator

import (
	"fmt"
	"regexp"
)

type IUserValidate interface {
	ValidatePhoneNumber(string) error
}

const (
	regexPhoneNumber string = `^(09|01[2|6|8|9])+([0-9]{8})$`
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
