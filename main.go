package main

import (
	"fmt"
	"log"
	"strconv"

	"tbox_backend/src/handlers"
	"tbox_backend/src/repository"
	"tbox_backend/src/routers"
	"tbox_backend/src/services"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	mongoURI := viper.GetString(`mongo.uri`)
	port := viper.GetString(`server.address`)
	redisURI := viper.GetString(`redis.uri`)
	redisConnectionPool := viper.GetString(`redis.connection_pool`)
	maxIdle, err := strconv.Atoi(redisConnectionPool)
	if err != nil {
		log.Fatal(err.Error())
	}

	session, err := mgo.Dial(mongoURI)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	p := &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   100,
		IdleTimeout: 5000,
	}
	p.Dial = func() (redis.Conn, error) {
		return redis.DialURL(redisURI)
	}

	router := gin.Default()

	userRepo := repository.NewUserRepository(session, p)
	userTokenRepo := repository.NewUserTokenRepository(session, p)
	userService := services.NewUserService(userRepo, userTokenRepo)
	userHandler := handlers.NewUserHandler(userService)

	// register routers
	routers.IndexRouter(router, userHandler)

	router.Run(port)
}
