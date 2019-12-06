package routers

import (
	"net/http"
	"tbox_backend/src/handlers"

	"github.com/gin-gonic/gin"
)

func IndexRouter(rg *gin.Engine, userHandler *handlers.UserHandler) {
	gr := rg.Group("/api/v1")
	{
		gr.GET("/ping", pingHandler)
		gr.POST("/login", userHandler.Login)
		gr.POST("/verify", userHandler.VerifyPhoneNumber)
		gr.POST("/resend-otp", userHandler.ResendOTP)
	}
}

// @Summary Ping server
// @Description Ping health check server
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"pong"
// @Router /api/v1/ping [get]
func pingHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong\n")
}
