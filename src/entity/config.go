package entity

type TwilioConfig struct {
	AccountSID  string `json:"account_sid"`
	AuthToken   string `json:"auth_token"`
	PhoneNumber string `json:"phone_nunber"`
}

type JWTConfig struct {
	SecretKey string `json:"secret_key"`
}
