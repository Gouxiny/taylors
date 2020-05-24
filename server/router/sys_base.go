package router

import (
	"github.com/gin-gonic/gin"
	"taylors/api/v1"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("register", v1.Register)
		BaseRouter.POST("login", v1.Login)
		BaseRouter.POST("captcha", v1.Captcha)
		BaseRouter.GET("captcha/:captchaId", v1.CaptchaImg)
	}
	return BaseRouter
}
