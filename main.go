package main

import (
	"fmt"
	"log"
	"strconv"

	"tbox_backend/src/entity"
	"tbox_backend/src/handlers"
	"tbox_backend/src/helper"
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

	// get config twilio
	twilioConfig := entity.TwilioConfig{
		AccountSID:  viper.GetString(`twilio.account_sid`),
		AuthToken:   viper.GetString(`twilio.auth_token`),
		PhoneNumber: viper.GetString(`twilio.phone_nunber`),
	}

	session, err := mgo.Dial(mongoURI)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   100,
		IdleTimeout: 5000,
	}
	pool.Dial = func() (redis.Conn, error) {
		return redis.DialURL(redisURI)
	}

	router := gin.Default()

	// import repository
	userRepo := repository.NewUserRepository(session, pool)
	userTokenRepo := repository.NewUserTokenRepository(session, pool)
	userOTPRepo := repository.NewUserOTPRepository(session, pool)

	// import helper
	twiliHelper := helper.NewTwilioHelper(twilioConfig)

	// import services
	userService := services.NewUserService(userRepo, userTokenRepo, userOTPRepo, twiliHelper)
	userHandler := handlers.NewUserHandler(userService)

	// register routers
	routers.IndexRouter(router, userHandler)

	router.Run(port)
}
