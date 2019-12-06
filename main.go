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
	// "github.com/go-redis/redis"
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

	// Configure list of places to look for IP address.
	// By default it's: "RemoteAddr", "X-Forwarded-For", "X-Real-IP"
	// If your application is behind a proxy, set "X-Forwarded-For" first.
	// limiter.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"})

	// // Limit only GET and POST requests.
	limiter.SetMethods([]string{"GET", "POST"})

	// // You can remove all entries at once.
	// limiter.RemoveHeader("X-Access-Token")

	// Or remove specific ones.
	// limiter.RemoveHeaderEntries("X-Access-Token", []string{"limitless-token"})

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
	twiliHelper := helper.NewTwilioHelper(twilioConfig)

	// import validate
	userValidator := validator.NewUserValidator()

	// import services
	userService := services.NewUserService(userRepo, userTokenRepo, userOTPRepo, twiliHelper, userValidator)
	userHandler := handlers.NewUserHandler(userService)

	router.Use(tollbooth_gin.LimitHandler(limiter))

	// register routers
	routers.IndexRouter(router, userHandler)

	router.Run(port)
}
