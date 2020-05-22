package router

import (
	v2 "gin-vue-admin/api/v2"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitStockRouter(Router *gin.RouterGroup) {
	StockRouter := Router.Group("stock").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		StockRouter.POST("createStock", v2.CreateStock)   // 新建Stock
		StockRouter.DELETE("deleteStock", v2.DeleteStock) //删除Stock
		StockRouter.PUT("updateStock", v2.UpdateStock)    //更新Stock
		StockRouter.GET("findStock", v2.FindStock)        // 根据ID获取Stock
		StockRouter.GET("getStockList", v2.GetStockList)  //获取Stock列表
	}
}
