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

type stockTop int

// @Tags Stock
// @Summary 获取Top列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.StockTopListReq true "获取Top列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stock/top/list [post]
func (*stockTop) StockTopList(c *gin.Context) {
	var req request.StockTopListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(fmt.Sprintf("参数错误，%v", err), c)
		return
	}

	filter := param.TopListParam{
		Name:             req.Name,
		Code:             req.Code,
		MarketCapitalMax: int64(req.MarketCapitalMax * 100000000),
		MarketCapitalMin: int64(req.MarketCapitalMin * 100000000),
		PercentMax:       req.PercentMax,
		PercentMin:       req.PercentMin,
		VolumeRatioMax:   req.VolumeRatioMax,
		VolumeRatioMin:   req.VolumeRatioMin,
		CurrentMax:       req.CurrentMax,
		CurrentMin:       req.CurrentMin,
	}

	stockList, err := service.StockTopService.TopList(filter)
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
