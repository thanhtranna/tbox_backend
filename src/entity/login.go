package entity

type Login struct {
	PhoneNumber string `json:"phone_number"`
}

type VerifyPhoneNumber struct {
	UserID string `json:"user_id"`
	OTP    string `json:"otp"`
}
