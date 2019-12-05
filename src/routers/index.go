package routers

import (
	"tbox_backend/src/handlers"

	"github.com/gin-gonic/gin"
)

func IndexRouter(rg *gin.Engine, userHandler *handlers.UserHandler) {
	gr := rg.Group("/v1")
	{
		gr.GET("/ping", pingHandler)
		// var userHandler handlers.NewUserHandler()
		gr.POST("/login", userHandler.Login)
	}
}

func pingHandler(ctx *gin.Context) {
	ctx.String(200, "pong\n")
}
