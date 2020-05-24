package service

import (
	"context"
	"gin-vue-admin/model"
	"gin-vue-admin/module/crawler"
	"taylors_proto/taylors_stock"
)

type stockCommonService struct {
}

func (*stockCommonService) HistoryList(code string) (stockList []model.Stock, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), _OverTime)
	defer cancel()

	req := &taylors_stock.HistoryListReq{
		Code: code,
	}
	historyListRsp, err := crawler.Grpc_cli.HistoryList(ctx, req)
	if err != nil {
		return
	}

	stockList = Conv(historyListRsp.StockList)
	return
}
