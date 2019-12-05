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
	uuid "github.com/satori/go.uuid"
)

type IUserRepo interface {
	FindByPhoneNumber(string) (entity.User, error)
}

type userRepository struct {
	mgo dl.IMgoDataLayer
	rds dl.IRedisDataLayer
}

func NewUserRepository(s *mgo.Session, rds *redis.Pool) IUserRepo {
	return &userRepository{
		mgo: dl.NewMgoDataLayer(s, common.CollectionUser),
		rds: dl.NewRedisDataLayer(rds),
	}
}

func (u *userRepository) FindByPhoneNumber(input string) (result entity.User, err error) {
	result, err = u.GetUserFromRedis(input)
	if err == nil && result.ID != "" {
		return result, nil
	}
	conditions := bson.D{
		{"phone_number", input},
	}

	err = u.mgo.FindOne(conditions, &result)
	if err != nil && err != mgo.ErrNotFound {
		return result, err
	}

	// if not exist in DB -> insert into DB
	if err == mgo.ErrNotFound {
		uuid := uuid.NewV4()
		result.ID = uuid.String()
		result.PhoneNumber = input
		result.IsVerify = false
		result.CreatedAt = common.GetVietNamTime()
		result.UpdatedAt = common.GetVietNamTime()

		err = u.mgo.Insert(result)
		if err != nil {
			return result, err
		}
		return result, u.CacheUserToRedis(result)
	}

	return result, u.CacheUserToRedis(result)
}

func (u *userRepository) CacheUserToRedis(input entity.User) error {
	dataRedis, err := json.Marshal(input)
	if err != nil {
		return err
	}
	return u.rds.SetString(u.getKeyRedisUser(input.PhoneNumber), -1, string(dataRedis))
}

func (u *userRepository) getKeyRedisUser(phoneNumber string) string {
	return fmt.Sprintf(common.KeyRedisUser, phoneNumber)
}

func (u *userRepository) GetUserFromRedis(input string) (result entity.User, err error) {
	dataRedis, err := u.rds.GetString(u.getKeyRedisUser(input))
	if err != nil && err.Error() != "redigo: nil returned" {
		return result, err
	}
	err = json.Unmarshal([]byte(dataRedis), &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
