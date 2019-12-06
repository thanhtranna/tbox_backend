package main

import (
	"fmt"
	"log"
	"time"

	"tbox_backend/src/entity"
	"tbox_backend/src/handlers"
	"tbox_backend/src/helper"
	"tbox_backend/src/repository"
	"tbox_backend/src/routers"
	"tbox_backend/src/services"
	"tbox_backend/src/validator"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "tbox_backend/docs"
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

// @title TBOX Backend API
// @version 1.0
// @description Swagger API for TBOX Backend.

// @BasePath /api/v1

func main() {
	mongoURI := viper.GetString(`mongo.uri`)
	port := viper.GetString(`server.address`)
	redisURI := viper.GetString(`redis.uri`)
	redisConnectionPool := viper.GetInt(`redis.connection_pool`)

	// get config twilio
	twilioConfig := entity.TwilioConfig{
		AccountSID:  viper.GetString(`twilio.account_sid`),
		AuthToken:   viper.GetString(`twilio.auth_token`),
		PhoneNumber: viper.GetString(`twilio.phone_nunber`),
	}

	// get config rate limit
	rateLimit := viper.GetFloat64(`rate_limit.limit`)
	// Create a limiter struct.
	limiter := tollbooth.NewLimiter(rateLimit, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Minute})

	// Limit only POST requests.
	limiter.SetMethods([]string{"POST"})

	authConfig := entity.JWTConfig{
		SecretKey: viper.GetString(`jwt.secret_key`),
	}

	session, err := mgo.Dial(mongoURI)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := &redis.Pool{
		MaxIdle:     redisConnectionPool,
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
	twilioHelper := helper.NewTwilioHelper(twilioConfig)
	authHelper := helper.NewAuthHelper(authConfig)

	// import validate
	userValidator := validator.NewUserValidator()

	// import services
	userService := services.NewUserService(userRepo, userTokenRepo, userOTPRepo, twilioHelper, authHelper, userValidator)
	userHandler := handlers.NewUserHandler(userService)

	router.Use(tollbooth_gin.LimitHandler(limiter))

	// register routers
	routers.IndexRouter(router, userHandler)
	// setup swagger
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run(port)
}
