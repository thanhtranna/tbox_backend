package common

const (
	KeyRedisUser      string = "tbox_backend_%s"
	KeyRedisUserToken string = "tbox_backend_user_token_%s"
	KeyRedisUserOTP   string = "tbox_backend_user_otp_%s"

	// Key check logic
	KeyRedisUserSendOTP string = "tbox_backend_sended_otp_%s"
)
