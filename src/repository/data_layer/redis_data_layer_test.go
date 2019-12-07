package data_layer_test

import (
	"testing"

	dl "tbox_backend/src/repository/data_layer"

	"github.com/gomodule/redigo/redis"
	"github.com/rafaeljusto/redigomock"
)

func TestSetStringNotExpire(t *testing.T) {
	conn := redigomock.NewConn()
	pool := &redis.Pool{
		// Return the same connection mock for each Get() call.
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}
	cmd := conn.Command("SET", "key", "value").ExpectMap(nil)

	redisCon := dl.NewRedisDataLayer(pool)

	err := redisCon.SetString("key", -1, "value")
	if err != nil {
		t.Fatal(err)
	}

	if conn.Stats(cmd) != 1 {
		t.Fatal("Command was not called!")
	}
}

func TestGetString(t *testing.T) {
	conn := redigomock.NewConn()
	pool := &redis.Pool{
		// Return the same connection mock for each Get() call.
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}
	cmd := conn.Command("GET", "key").Expect("value")

	redisCon := dl.NewRedisDataLayer(pool)

	dataStr, err := redisCon.GetString("key")
	if err != nil {
		t.Fatal(err)
	}

	if conn.Stats(cmd) != 1 {
		t.Fatal("Command was not called!")
	}

	if dataStr != "value" {
		t.Fatalf("Invalid value. Expected 'value' and got '%s'\n", dataStr)
	}
}

func TestDeleteString(t *testing.T) {
	conn := redigomock.NewConn()
	pool := &redis.Pool{
		// Return the same connection mock for each Get() call.
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}
	cmd := conn.Command("DEL", "key").Expect(nil)

	redisCon := dl.NewRedisDataLayer(pool)

	err := redisCon.DeleteKey("key")
	if err != nil {
		t.Fatal(err)
	}

	if conn.Stats(cmd) != 1 {
		t.Fatal("Command was not called!")
	}
}
