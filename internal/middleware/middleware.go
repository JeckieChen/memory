package middleware

import (
	"memory/internal/service"

	"github.com/gin-gonic/gin"
)

func NewCasbinAuth(srv *service.CasbinService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := srv.Enforcer.LoadPolicy()
		if err != nil {
			ctx.String(500, err.Error())
			ctx.Abort()
			return
		}
		// 简便起见，假设用户从url传递 /xxxx?username=leo，实际应用可以结合jwt等鉴权
		username, _ := ctx.GetQuery("username")
		// jwtToken := ctx.GetHeader("token")

		ok, err := srv.Enforcer.Enforce(username, ctx.Request.URL.Path, ctx.Request.Method)
		if err != nil {
			ctx.String(500, err.Error())
			ctx.Abort()
			return
		} else if !ok {
			ctx.String(403, "验证权限失败!")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
