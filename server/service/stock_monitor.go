package service

import (
	"context"
	"strings"
	"taylors/dao"
	"taylors/model"
	"taylors/model/request"
	"taylors/model/response"
	"taylors/module/crawler"
	"taylors/utils"
	"taylors_proto/taylors_stock"
	"time"
)

type stockMonitorService struct {
}

func (*stockMonitorService) MonitorOne(id int64, userId uint) (stockMonitor *model.StockMonitor, err error) {
	stockMonitor, err = dao.StockMonitorDao.FindByUserAndId(userId, id)
	if err != nil {
		return
	}
	return
}

func (srv *stockMonitorService) MonitorList(uid uint, filter request.MonitorListReq) (stockList []response.StockMonitorModel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), _OverTime)
	defer cancel()

	monitorList := []*model.StockMonitor{}
	if !filter.IsDay {
		monitorList, err = dao.StockMonitorDao.ListByUserNotDay(uid)
		if err != nil {
			return
		}
	} else {
		monitorList, err = dao.StockMonitorDao.ListByUserAndDay(uid, utils.NowUnix())
		if err != nil {
			return
		}
	}

	monitorMap := make(map[string]*model.StockMonitor, 0)
	for _, monitor := range monitorList {
		monitorMap[monitor.Symbol] = monitor
	}

	req := &taylors_stock.MonitorReq{Symbol: []string{}}
	monitorRsp, err := crawler.Grpc_cli.Monitor(ctx, req)
	if err != nil {
		return
	}
	stocks := make([]response.StockMonitorModel, 0)
	stockConvList := Conv(monitorRsp.StockList)
	for _, stock := range stockConvList {
		symbolAndName := strings.Split(stock.Symbol, "-")
		if symbolAndName == nil || len(symbolAndName) == 0 {
			continue
		}
		stockMonitor, ok := monitorMap[symbolAndName[0]]
		if ok {
			stock.Id = stockMonitor.Id
			stockModel := ConvStockMonitorModel(stock)
			stockModel.MonitorHigh = stockMonitor.MonitorHigh
			stockModel.MonitorLow = stockMonitor.MonitorLow
			stocks = append(stocks, stockModel)
		}
	}

	for _, stock := range stocks {
		if filter.Name != "" {
			if !strings.Contains(stock.Name, filter.Name) {
				continue
			}
		}

		if filter.Symbol != "" {
			if !strings.Contains(stock.Symbol, filter.Symbol) {
				continue
			}
		}

		if filter.MarketCapitalMax > 0 {
			if stock.MarketCapital > filter.MarketCapitalMax {
				continue
			}
		}
		if filter.MarketCapitalMin > 0 {
			if stock.MarketCapital < filter.MarketCapitalMin {
				continue
			}
		}
		if filter.PercentMax > 0 {
			if stock.Percent > filter.PercentMax {
				continue
			}
		}
		if filter.PercentMin > 0 {
			if stock.Percent < filter.PercentMin {
				continue
			}
		}
		if filter.VolumeRatioMax > 0 {
			if stock.VolumeRatio > filter.VolumeRatioMax {
				continue
			}
		}
		if filter.VolumeRatioMin > 0 {
			if stock.VolumeRatio < filter.VolumeRatioMin {
				continue
			}
		}
		if filter.CurrentMax > 0 {
			if stock.Current > filter.CurrentMax {
				continue
			}
		}
		if filter.CurrentMin > 0 {
			if stock.Current < filter.CurrentMin {
				continue
			}
		}
		stockList = append(stockList, stock)
	}

	return
}

func (srv *stockMonitorService) AddMonitor(isDay bool, symbol string, monitorHigh, monitorLow float64, userId uint) (err error) {
	monitor := model.StockMonitor{
		Symbol:      symbol,
		MonitorHigh: monitorHigh,
		MonitorLow:  monitorLow,
		UserId:      userId,
		IsDay:       isDay,
		CreateTime:  time.Now().Unix(),
	}
	if isDay {
		monitor.Day = utils.NowUnix()
	}

	err = dao.StockMonitorDao.Save(monitor)
	go srv.SyncMonitor()

	return
}

func (srv *stockMonitorService) DelMonitor(id int64, userId uint) (err error) {
	err = dao.StockMonitorDao.DelById(id, userId)
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
