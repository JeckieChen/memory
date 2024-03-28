package controller

import (
	"memory/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (ctrl *UserController) Login(ctx *gin.Context) {
	var p service.LoginInfo
	ctx.BindJSON(&p)
	token, err := service.Login(&p)
	if err != nil {
		ctx.String(500, "登录失败"+err.Error())
	} else {
		ctx.JSON(200, token)
	}
}

func (ctrl *UserController) Register(ctx *gin.Context) {
	var p service.UserInfo
	ctx.BindJSON(&p)
	status, err := service.Register(&p)
	if status == false {
		ctx.String(500, "登录失败"+err.Error())
	} else {
		ctx.JSON(200, "注册成功")
	}
}
