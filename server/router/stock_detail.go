package router

import (
	"github.com/gin-gonic/gin"
	v1 "taylors/api/v1"
	"taylors/middleware"
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
		StockRouter.POST("monitor/one", v1.StockMonitorCon.StockMonitorOne)
		StockRouter.POST("monitor/list", v1.StockMonitorCon.StockMonitorList)
		StockRouter.POST("monitor/add", v1.StockMonitorCon.AddMonitor)
		StockRouter.POST("monitor/del", v1.StockMonitorCon.DelMonitor)
		StockRouter.POST("monitor/update", v1.StockMonitorCon.UpdateMonitor)
	}
}
