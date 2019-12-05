package data_layer

import (
	"github.com/gomodule/redigo/redis"
)

type IRedisDataLayer interface {
	SetHSET(key, field, value string)
	IsNilErr(err error) bool
	GetHSET(key, field string) (string, error)
	DeleteKeyHSET(key, field string) error
	GetString(key string) (string, error)
	SetString(key string, expire int, value string) error
	DeleteKey(key string) error
}

type redisDataLayer struct {
	connection *redis.Pool
}

func NewRedisDataLayer(connection *redis.Pool) IRedisDataLayer {
	return &redisDataLayer{
		connection: connection,
	}
}

func (r *redisDataLayer) IsNilErr(err error) bool {
	return err == redis.ErrNil
}

func (r *redisDataLayer) getConnect() redis.Conn {
	return r.connection.Get()
}

func (m *redisDataLayer) SetHSET(key, field, value string) {
	c := m.getConnect()
	defer c.Close()
	c.Do("HSET", key, field, value)
}

func (m *redisDataLayer) GetHSET(key, field string) (string, error) {
	c := m.getConnect()
	defer c.Close()
	s, err := redis.String(c.Do("HGET", key, field))
	return s, err
}

func (m *redisDataLayer) DeleteKeyHSET(key, field string) error {
	c := m.getConnect()
	defer c.Close()

	_, err := c.Do("HDEL", key, field)
	return err
}

func (m *redisDataLayer) GetString(key string) (string, error) {
	c := m.getConnect()
	defer c.Close()

	s, err := redis.String(c.Do("GET", key))
	return s, err
}

func (m *redisDataLayer) GetInt(key string) (int, error) {
	c := m.getConnect()
	defer c.Close()

	return redis.Int(c.Do("GET", key))
}

func (m *redisDataLayer) SetString(key string, expire int, value string) error {
	c := m.getConnect()
	defer c.Close()

	var err error
	if expire <= 0 {
		_, err = c.Do("SET", key, value)
	} else {
		_, err = c.Do("SETEX", key, expire, value)
	}
	return err
}

func (m *redisDataLayer) DeleteKey(key string) error {
	c := m.getConnect()
	defer c.Close()

	_, err := c.Do("DEL", key)
	return err
}
