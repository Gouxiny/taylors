package service

import (
	"context"
	"taylors/model"
	"taylors/module/crawler"
	"taylors_proto/taylors_stock"
)

type stockMonitorService struct {
}

func (*stockMonitorService) MonitorList(symbolList []string) (stockList []model.Stock, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), _OverTime)
	defer cancel()

	req := &taylors_stock.MonitorReq{Symbol: symbolList}
	monitorRsp, err := crawler.Grpc_cli.Monitor(ctx, req)
	if err != nil {
		return
	}

	stockList = Conv(monitorRsp.StockList)
	return
}
