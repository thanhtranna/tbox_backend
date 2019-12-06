package entity

type User struct {
	ID          string `json:"user_id" bson:"user_id"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
	IsVerify    bool   `json:"is_verify" bson:"is_verify"`
	CreatedAt   int64  `json:"created_at" bson:"created_at"`
	UpdatedAt   int64  `json:"updated_at" bson:"updated_at"`
}

type UserToken struct {
	ID        string `json:"user_id" bson:"user_id"`
	Token     string `json:"token" bson:"token"`
	CreatedAt int64  `json:"created_at" bson:"created_at"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at"`
}

type UserOTP struct {
	ID         string `json:"user_id" bson:"user_id"`
	OTP        string `json:"otp" bson:"otp"`
	TimeExpire int64  `json:"time_expire" bson:"time_expire"`
	CreatedAt  int64  `json:"created_at" bson:"created_at"`
	UpdatedAt  int64  `json:"updated_at" bson:"updated_at"`
}
