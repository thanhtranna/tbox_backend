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

type IUserTokenRepo interface {
	GetTokenByUserID(string) (entity.UserToken, error)
	Save(data entity.UserToken) error
}

type userTokenRepository struct {
	mgo dl.IMgoDataLayer
	rds dl.IRedisDataLayer
}

func NewUserTokenRepository(s *mgo.Session, rds *redis.Pool) IUserTokenRepo {
	return &userTokenRepository{
		mgo: dl.NewMgoDataLayer(s, common.CollectionUserToken),
		rds: dl.NewRedisDataLayer(rds),
	}
}

func (u *userTokenRepository) GetTokenByUserID(userID string) (result entity.UserToken, err error) {
	result, err = u.GetUserTokenFromRedis(userID)
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
	return result, u.CacheUserTokenToRedis(result)
}

func (u *userTokenRepository) Save(data entity.UserToken) error {
	return u.mgo.Insert(data)
}

func (u *userTokenRepository) getKeyRedisUserToken(input string) string {
	return fmt.Sprintf(common.KeyRedisUserToken, input)
}

func (u *userTokenRepository) CacheUserTokenToRedis(input entity.UserToken) error {
	dataRedis, err := json.Marshal(input)
	if err != nil {
		return err
	}
	return u.rds.SetString(u.getKeyRedisUserToken(input.ID), -1, string(dataRedis))
}

func (u *userTokenRepository) GetUserTokenFromRedis(input string) (result entity.UserToken, err error) {
	dataRedis, err := u.rds.GetString(u.getKeyRedisUserToken(input))
	if err != nil && err.Error() != "redigo: nil returned" {
		return result, err
	}
	err = json.Unmarshal([]byte(dataRedis), &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
