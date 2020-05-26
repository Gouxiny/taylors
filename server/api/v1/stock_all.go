package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"taylors/global/response"
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
func (*stockAll) StockAllDetailList(c *gin.Context) {
	var req request.AllDetailListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(fmt.Sprintf("参数错误，%v", err), c)
		return
	}

	stockList, err := service.StockAllService.AllDetailList()
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
