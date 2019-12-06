package data_layer

import (
	"github.com/gomodule/redigo/redis"
)

type IRedisDataLayer interface {
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

func (r *redisDataLayer) getConnect() redis.Conn {
	return r.connection.Get()
}

func (m *redisDataLayer) GetString(key string) (string, error) {
	c := m.getConnect()
	defer c.Close()

	s, err := redis.String(c.Do("GET", key))
	return s, err
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
