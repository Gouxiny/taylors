package service

import (
	"context"
	"gin-vue-admin/model"
	"gin-vue-admin/module/crawler"
	"taylors_proto/taylors_stock"
)

type stockAllService struct {
}

func (*stockAllService) AllDetailList(pageNum, pageSize int) (stockList []model.Stock, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), _OverTime)
	defer cancel()
	req := &taylors_stock.AllDetailReq{}
	allListRsp, err := crawler.Grpc_cli.AllDetail(ctx, req)
	if err != nil {
		return
	}
	stockList = Conv(allListRsp.StockList)
	return
}
