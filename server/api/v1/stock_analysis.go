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

type stockAnalysis int

// @Tags Stock
// @Summary 获取数据分析列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AnalysisListReq true "获取数据分析列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stock/analysis/list [post]
func (*stockAnalysis) StockAnalysisList(c *gin.Context) {
	var req request.AnalysisListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(fmt.Sprintf("参数错误，%v", err), c)
		return
	}

	param := &param.AnalysisListParam{
		Page:             req.Page,
		PageSize:         req.PageSize,
		MarketCapitalMax: int64(req.MarketCapitalMax * 100000000),
		MarketCapitalMin: int64(req.MarketCapitalMin * 100000000),
		PercentMax:       req.PercentMax,
		PercentMin:       req.PercentMin,
		VolumeRatioMax:   req.VolumeRatioMax,
		VolumeRatioMin:   req.VolumeRatioMin,
		CurrentMax:       req.CurrentMax,
		CurrentMin:       req.CurrentMin,
		StartTime:        req.StartTime,
		EndTime:          req.EndTime,
		DayMin:           req.DayMin,
	}

	stockList, total, err := service.StockAnalysisService.AnalysisList(param)
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
