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

// @Summary Login with phone number
// @Description User can login in system
// @Accept  json
// @Produce  json
// @Param   phone_number	path	string	true	"0987654321"
// @Success 200 {object} common.DataFormat	"ok"
// @Failure 400 {object} common.DataFormat "something went wrong"
// @Failure 404 {object} common.DataFormat "Not Found"
// @Router /api/v1/login [post]
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

// @Summary Verify phone number of user
// @Description User verify phone number before accept system
// @Accept  json
// @Produce  json
// @Param 	user_id	path	string	true	"9c070de4-d6aa-4d53-b9f3-02eacb5a1d05"
// @Param   otp     path    string	true	"123456"
// @Success 200 {object} common.DataFormat "ok"
// @Failure 400 {object} common.DataFormat "something went wrong"
// @Failure 404 {object} common.DataFormat "Not Found"
// @Router /api/v1/verify [post]
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

// @Summary Resend OTP to phone number of user
// @Description User resend otp if user not receive otp before
// @Accept  json
// @Produce  json
// @Param   user_id     path    string     true        "9c070de4-d6aa-4d53-b9f3-02eacb5a1d05"
// @Success 200 {object} common.DataFormat "ok"
// @Failure 400 {object} common.DataFormat "something went wrong"
// @Failure 404 {object} common.DataFormat "Not Found"
// @Router /api/v1/resend-otp [post]
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
