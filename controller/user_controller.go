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
	service.Login(&p)
}
