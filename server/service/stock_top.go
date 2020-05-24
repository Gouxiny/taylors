package service

import (
	"context"
	"gin-vue-admin/model"
	"gin-vue-admin/module/crawler"
	"taylors_proto/taylors_stock"
)

type stockTopService struct {
}

func (*stockTopService) TopList() (stockList []model.Stock, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), _OverTime)
	defer cancel()
	req := &taylors_stock.TopReq{}
	topRsp, err := crawler.Grpc_cli.ListTop(ctx, req)
	if err != nil {
		return
	}
	stockList = Conv(topRsp.StockList)
	return
}
