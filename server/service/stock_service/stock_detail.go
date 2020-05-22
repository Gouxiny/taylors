package stock_service

import (
	"context"
	"gin-vue-admin/crawler_client"
	"gin-vue-admin/model"
	"taylors_proto/taylors_stock"
	"time"
)

var _OverTime = time.Second * 10

type stockService struct {
}

func (*stockService) TopList() (stockList []model.Stock, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), _OverTime)
	defer cancel()
	req := &taylors_stock.TopReq{}
	topRsp, err := crawler_client.Grpc_cli.ListTop(ctx, req)
	if err != nil {
		return
	}
	stockList = Conv(topRsp.StockList)
	return
}

func (*stockService) AllDetailList() (stockList []model.Stock, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), _OverTime)
	defer cancel()
	req := &taylors_stock.AllDetailReq{}
	allListRsp, err := crawler_client.Grpc_cli.AllDetail(ctx, req)
	if err != nil {
		return
	}
	stockList = Conv(allListRsp.StockList)
	return
}

func (*stockService) MonitorList(symbolList []string) (stockList []model.Stock, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), _OverTime)
	defer cancel()

	req := &taylors_stock.MonitorReq{Symbol: symbolList}
	monitorRsp, err := crawler_client.Grpc_cli.Monitor(ctx, req)
	if err != nil {
		return
	}

	stockList = Conv(monitorRsp.StockList)
	return
}

func (*stockService) HistoryList(code string) (stockList []model.Stock, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), _OverTime)
	defer cancel()

	req := &taylors_stock.HistoryListReq{
		Code: code,
	}
	historyListRsp, err := crawler_client.Grpc_cli.HistoryList(ctx, req)
	if err != nil {
		return
	}

	stockList = Conv(historyListRsp.StockList)
	return
}
