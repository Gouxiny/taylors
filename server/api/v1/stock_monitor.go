package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"taylors/global/response"
	"taylors/model/request"
	resp "taylors/model/response"
	"taylors/service"
)

type stockMonitor int

// @Tags StockMonitor
// @Summary 获取监控
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.MonitorOneReq true "获取监控"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stock/monitor/one [post]
func (*stockMonitor) StockMonitorOne(c *gin.Context) {
	var req request.MonitorOneReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(fmt.Sprintf("参数错误，%v", err), c)
		return
	}
	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	stockMonitor, err := service.StockMonitorService.MonitorOne(req.Id, waitUse.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败，%v", err), c)
	} else {
		response.OkWithData(stockMonitor, c)
	}
}

// @Tags StockMonitor
// @Summary 获取监控列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stock/monitor/list [post]
func (*stockMonitor) StockMonitorList(c *gin.Context) {
	var req request.MonitorListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(fmt.Sprintf("参数错误，%v", err), c)
		return
	}
	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	stockList, err := service.StockMonitorService.MonitorList(waitUse.ID, req)
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

// @Tags StockMonitor
// @Summary 添加监控
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AddMonitorReq true "添加监控"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加监控成功"}"
// @Router /stock/monitor/add [post]
func (*stockMonitor) AddMonitor(c *gin.Context) {
	var req request.AddMonitorReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(fmt.Sprintf("参数错误，%v", err), c)
		return
	}
	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	err := service.StockMonitorService.AddMonitor(req.IsDay, req.Code, req.MonitorHigh, req.MonitorLow, waitUse.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags StockMonitor
// @Summary 删除监控
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DelMonitorReq true "删除监控"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除监控成功"}"
// @Router /stock/monitor/del [post]
func (*stockMonitor) DelMonitor(c *gin.Context) {
	var req request.DelMonitorReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(fmt.Sprintf("参数错误，%v", err), c)
		return
	}
	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	err := service.StockMonitorService.DelMonitor(req.Id, waitUse.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags StockMonitor
// @Summary 更新监控
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UpdateMonitorReq true "更新监控"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新监控成功"}"
// @Router /stock/monitor/update [post]
func (*stockMonitor) UpdateMonitor(c *gin.Context) {
	var req request.UpdateMonitorReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(fmt.Sprintf("参数错误，%v", err), c)
		return
	}
	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	err := service.StockMonitorService.UpdateMonitor(req.MonitorHigh, req.MonitorLow, req.Id, waitUse.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
