package validator

import (
	"fmt"
	"regexp"
	"tbox_backend/src/entity"
)

type IUserValidate interface {
	ValidatePhoneNumber(string) error
	ValidateUserID(string) error
	ValidateOTP(string) error
	ValidateVerifyPhoneNumber(entity.VerifyPhoneNumber) error
}

const (
	regexOTP         string = `^[0-9]{6}$`
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

func (u *userValidator) ValidateOTP(otp string) error {
	regex := regexp.MustCompile(regexOTP)
	if !regex.MatchString(otp) {
		return fmt.Errorf("otp invalid")
	}

	return nil
}

func (u *userValidator) ValidateVerifyPhoneNumber(input entity.VerifyPhoneNumber) error {
	if err := u.ValidateUserID(input.UserID); err != nil {
		return err
	}
	if err := u.ValidateOTP(input.OTP); err != nil {
		return err
	}

	return nil
}
