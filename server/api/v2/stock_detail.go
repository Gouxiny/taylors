package v2

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service/stock_service"
	"github.com/gin-gonic/gin"
)

// @Tags Stock
// @Summary 获取Top列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.StockTopListReq true "获取Top列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stock/stockTopList [post]
func StockTopList(c *gin.Context) {
	var req request.StockTopListReq
	_ = c.ShouldBindJSON(&req)
	stockList, err := stock_service.StockService.TopList()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取Top失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     stockList,
			Total:    len(stockList),
			Page:     1,
			PageSize: len(stockList),
		}, c)
	}
}

// @Tags Stock
// @Summary 获取监控列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stock true "获取监控列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stock/stockMonitorList [post]
func StockMonitorList(c *gin.Context) {
	var req request.MonitorListReq
	_ = c.ShouldBindJSON(&req)
	stockList, err := stock_service.StockService.MonitorList(req.SymbolList)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     stockList,
			Total:    len(stockList),
			Page:     1,
			PageSize: len(stockList),
		}, c)
	}
}

// @Tags Stock
// @Summary 获取监控列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stock true "获取监控列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stock/stockMonitorList [post]
func StockAllDetailList(c *gin.Context) {
	var req request.MonitorListReq
	_ = c.ShouldBindJSON(&req)
	stockList, err := stock_service.StockService.MonitorList(req.SymbolList)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     stockList,
			Total:    len(stockList),
			Page:     1,
			PageSize: len(stockList),
		}, c)
	}
}
