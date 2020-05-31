package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"taylors/global/response"
	"taylors/model/param"
	"taylors/model/request"
	resp "taylors/model/response"
	"taylors/service"
)

type stockAll int

// @Tags Stock
// @Summary 获取所有列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AllDetailListReq true "获取所有列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stock/all/list [post]
func (*stockAll) StockAllList(c *gin.Context) {
	var req request.AllListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(fmt.Sprintf("参数错误，%v", err), c)
		return
	}

	param := &param.AllListParam{
		Page:             req.Page,
		PageSize:         req.PageSize,
		Name:             req.Name,
		Code:             req.Code,
		MarketCapitalMax: req.MarketCapitalMax * 100000000,
		MarketCapitalMin: req.MarketCapitalMin * 100000000,
		PercentMax:       req.PercentMax,
		PercentMin:       req.PercentMin,
		VolumeRatioMax:   req.VolumeRatioMax,
		VolumeRatioMin:   req.VolumeRatioMin,
		CurrentMax:       req.CurrentMax,
		CurrentMin:       req.CurrentMin,
	}

	stockList, total, err := service.StockAllService.AllList(param)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     stockList,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, c)
	}
}
