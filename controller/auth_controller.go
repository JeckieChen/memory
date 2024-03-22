package controller

import (
	"github.com/gin-gonic/gin"
)

type AuthController struct {
}

func (controller *AuthController) UpdatePolicy(ctx *gin.Context) {
	// var p service.RolePolicy
	// ctx.BindJSON(&p)
	// casbinService, _ := service.NewCasbinService()
	// err := casbinService.CreateRolePolicy(p)
	// if err != nil {
	// 	ctx.String(500, "创建角色策略失败: "+err.Error())
	// } else {
	// 	ctx.JSON(200, "成功!")
	// }
}
