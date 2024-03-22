package config

import (
	"memory/controller"
	"memory/service"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	routerAuth(router)
}

func routerAuth(engine *gin.Engine) {
	group := engine.Group("/")
	group.Use(NewCasbinAuth(&service.CasbinService{}))
	{
		con := &controller.AuthController{}
		group.POST("/casbin/rolepolicy", con.UpdatePolicy)
	}
}
