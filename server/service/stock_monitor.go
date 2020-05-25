package service

import (
	"context"
	"taylors/dao"
	"taylors/model"
	"taylors/module/crawler"
	"taylors/utils"
	"taylors_proto/taylors_stock"
	"time"
)

type stockMonitorService struct {
}

func (*stockMonitorService) MonitorOne(symbol string, userId uint) (stockMonitor *model.StockMonitor, err error) {
	stockMonitor, err = dao.StockMonitorDao.FindByUserAndSymbol(userId, symbol)
	if err != nil {
		return
	}
	return
}

func (*stockMonitorService) MonitorList() (stockList []model.Stock, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), _OverTime)
	defer cancel()

	req := &taylors_stock.MonitorReq{Symbol: []string{}}
	monitorRsp, err := crawler.Grpc_cli.Monitor(ctx, req)
	if err != nil {
		return
	}

	stockList = Conv(monitorRsp.StockList)
	return
}

func (srv *stockMonitorService) AddMonitor(symbol string, monitorHigh, monitorLow float64, userId uint) (err error) {
	monitor := model.StockMonitor{
		Symbol:      symbol,
		MonitorHigh: monitorHigh,
		MonitorLow:  monitorLow,
		UserId:      userId,
		CreateTime:  time.Now().Unix(),
	}
	err = dao.StockMonitorDao.Save(monitor)
	go srv.SyncMonitor()

	return
}

func (srv *stockMonitorService) DelMonitor(symbol string, userId uint) (err error) {
	err = dao.StockMonitorDao.DelBySymbol(symbol, userId)
	if err != nil {
		return
	}
	go srv.SyncMonitor()
	return
}

func (srv *stockMonitorService) UpdateMonitor(monitorHigh, monitorLow float64, id int64, uid uint) (err error) {
	err = dao.StockMonitorDao.UpdateByMonitorNum(monitorHigh, monitorLow, id, uid)
	if err != nil {
		return
	}
	go srv.SyncMonitor()
	return
}

func (*stockMonitorService) SyncMonitor() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), _OverTime)
	defer cancel()
	stockMonitorList, err := dao.StockMonitorDao.ListNotDel()
	if err != nil {
		return
	}

	symbolList := make([]string, 0)
	for _, stockMonitor := range stockMonitorList {
		symbolList = append(symbolList, stockMonitor.Symbol)
	}

	req := &taylors_stock.MonitorReq{Symbol: utils.SliceUniqueString(symbolList)}
	_, err = crawler.Grpc_cli.Monitor(ctx, req)

	return
}
