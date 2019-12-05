package validator

import (
	"fmt"
	"regexp"
)

type IUserValidate interface {
	ValidatePhoneNumber(string) error
}

const (
	regexPhoneNumber string = `^0(1\d{9}|9\d{8})$`
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
