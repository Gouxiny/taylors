package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitStockRouter(Router *gin.RouterGroup) {
	StockRouter := Router.Group("stock").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		StockRouter.POST("top/list", v1.StockTopCon.StockTopList)
	}

	{
		StockRouter.POST("all/list", v1.StockAllCon.StockAllDetailList)
	}

	{
		StockRouter.POST("monitor/list", v1.StockMonitorCon.StockMonitorList)
	}
}
