package handlers

import (
	"net/http"

	"tbox_backend/src/common"
	"tbox_backend/src/entity"
	"tbox_backend/src/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.IUserService
}

func NewUserHandler(userService services.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) Login(ctx *gin.Context) {
	var input entity.Login

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseFail(err.Error()))
		return
	}
	token, err := uh.userService.Login(input.PhoneNumber)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseFail(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, common.ResponseSuccess(token))
}

func (uh *UserHandler) VerifyPhoneNumber(ctx *gin.Context) {
	var input entity.VerifyPhoneNumber

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseFail(err.Error()))
		return
	}
	token, err := uh.userService.VerifyPhoneNumber(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseFail(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, common.ResponseSuccess(token))
}

func (uh *UserHandler) ResendOTP(ctx *gin.Context) {
	var input entity.ResendOTP

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseFail(err.Error()))
		return
	}
	err := uh.userService.ResendOTP(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseFail(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, common.ResponseSuccess("OK"))
}
