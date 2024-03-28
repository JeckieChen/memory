package config

import (
	"memory/controller"
	"memory/service"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	routerAuth(router)
	routerUser(router)
}

func routerAuth(engine *gin.Engine) {
	group := engine.Group("/")
	group.Use(NewCasbinAuth(service.CasbinSrv))
	{
		con := &controller.AuthController{}
		group.POST("/casbin/rolepolicy", con.UpdatePolicy)
		group.DELETE("/casbin/rolepolicy", con.DeletePolicy)
	}
}

func routerUser(engine *gin.Engine) {
	group := engine.Group("/user")
	{
		con := &controller.UserController{}
		group.POST("/login", con.Login)
	}
}
