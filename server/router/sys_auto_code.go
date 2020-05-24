package router

import (
	"github.com/gin-gonic/gin"
	"taylors/api/v1"
	"taylors/middleware"
)

func InitAutoCodeRouter(Router *gin.RouterGroup) {
	AutoCodeRouter := Router.Group("autoCode").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		AutoCodeRouter.POST("createTemp", v1.CreateTemp) //创建自动化代码
	}
}
