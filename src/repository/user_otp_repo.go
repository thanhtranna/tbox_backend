package repository

import (
	"encoding/json"
	"fmt"
	"tbox_backend/src/common"
	"tbox_backend/src/entity"
	dl "tbox_backend/src/repository/data_layer"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gomodule/redigo/redis"
)

type IUserOTPRepo interface {
	GetOTPByUserID(string) (entity.UserOTP, error)
	CheckSendedOTP(string) (string, bool)
	Save(entity.UserOTP) error
	CacheOTPWithUser(string, string) error
}

type userOTPRepository struct {
	mgo dl.IMgoDataLayer
	rds dl.IRedisDataLayer
}

func NewUserOTPRepository(s *mgo.Session, rds *redis.Pool) IUserOTPRepo {
	return &userOTPRepository{
		mgo: dl.NewMgoDataLayer(s, common.CollectionUserOtp),
		rds: dl.NewRedisDataLayer(rds),
	}
}

func (u *userOTPRepository) GetOTPByUserID(userID string) (result entity.UserOTP, err error) {
	result, err = u.GetUserOTPFromRedis(userID)
	if err == nil && result.ID != "" {
		return result, nil
	}
	conditions := bson.D{
		{"user_id", userID},
	}
	err = u.mgo.FindOne(conditions, &result)
	if err != nil && err != mgo.ErrNotFound {
		return result, err
	}
	return result, u.CacheUserOTPToRedis(result)
}

func (u *userOTPRepository) Save(data entity.UserOTP) error {
	conditions := bson.D{
		{"user_id", data.ID},
	}
	return u.mgo.Upsert(conditions, data)
}

func (u *userOTPRepository) CheckSendedOTP(userID string) (string, bool) {
	dataRedis, _ := u.rds.GetString(u.getKeyOTPOfUser(userID))
	if dataRedis != "" {
		return dataRedis, true
	}
	return "", false
}

// Key redis check OTP has sended
func (u *userOTPRepository) getKeyOTPOfUser(userID string) string {
	return fmt.Sprintf(common.KeyRedisUserSendOTP, userID)
}

// Key redis get data cache OTP of user
func (u *userOTPRepository) getKeyRedisUserOTP(input string) string {
	return fmt.Sprintf(common.KeyRedisUserOTP, input)
}

func (u *userOTPRepository) CacheUserOTPToRedis(input entity.UserOTP) error {
	dataRedis, err := json.Marshal(input)
	if err != nil {
		return err
	}
	return u.rds.SetString(u.getKeyRedisUserOTP(input.ID), -1, string(dataRedis))
}

func (u *userOTPRepository) GetUserOTPFromRedis(input string) (result entity.UserOTP, err error) {
	dataRedis, err := u.rds.GetString(u.getKeyRedisUserOTP(input))
	if err != nil && err != redis.ErrNil {
		return result, err
	}
	err = json.Unmarshal([]byte(dataRedis), &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (u *userOTPRepository) CacheOTPWithUser(userID string, OTP string) error {
	return u.rds.SetString(u.getKeyOTPOfUser(userID), common.TimeCacheOTP, OTP)
}
