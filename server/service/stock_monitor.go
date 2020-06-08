package service

import (
	"strings"
	"taylors/crawler"
	"taylors/dao"
	"taylors/model"
	"taylors/model/request"
	"taylors/model/response"
	"taylors/utils"
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

	monitorCodes := make([]string, 0)
	monitorMap := make(map[string]*model.StockMonitor, 0)
	for _, monitor := range monitorList {
		monitorCodes = append(monitorCodes, monitor.Code)
		monitorMap[monitor.Code] = monitor
	}

	marketList := crawler.NewDongFangCrawler().Monitor(monitorCodes)
	for _, stock := range marketList {
		if filter.Name != "" {
			if !strings.Contains(stock.Name, filter.Name) {
				continue
			}
		}

		if filter.Code != "" {
			if !strings.Contains(stock.Code, filter.Code) {
				continue
			}
		}

		if filter.MarketCapitalMax != 0 {
			if stock.MarketCapital > int64(filter.MarketCapitalMax) {
				continue
			}
		}
		if filter.MarketCapitalMin != 0 {
			if stock.MarketCapital < int64(filter.MarketCapitalMin) {
				continue
			}
		}
		if filter.PercentMax != 0 {
			if stock.Percent > filter.PercentMax {
				continue
			}
		}
		if filter.PercentMin != 0 {
			if stock.Percent < filter.PercentMin {
				continue
			}
		}
		if filter.VolumeRatioMax != 0 {
			if stock.VolumeRatio > filter.VolumeRatioMax {
				continue
			}
		}
		if filter.VolumeRatioMin != 0 {
			if stock.VolumeRatio < filter.VolumeRatioMin {
				continue
			}
		}
		if filter.CurrentMax != 0 {
			if stock.Current > filter.CurrentMax {
				continue
			}
		}
		if filter.CurrentMin != 0 {
			if stock.Current < filter.CurrentMin {
				continue
			}
		}

		stock.MarketCapital = stock.MarketCapital / 100000000

		monitorModel := response.StockMonitorModel{}
		monitorModel.Stock = stock
		stockPO, ok := monitorMap[stock.Code]
		if ok {
			monitorModel.MonitorLow = stockPO.MonitorLow
			monitorModel.MonitorHigh = stockPO.MonitorHigh
			monitorModel.Id = stockPO.Id
			stockList = append(stockList, monitorModel)
		}
	}
	return
}

func (srv *stockMonitorService) AddMonitor(isDay bool, code string, monitorHigh, monitorLow float64, userId uint) (err error) {
	monitor := model.StockMonitor{
		Code:        code,
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
	return
}

func (srv *stockMonitorService) DelMonitor(id int64, userId uint) (err error) {
	err = dao.StockMonitorDao.DelById(id, userId)
	if err != nil {
		return
	}
	return
}

func (srv *stockMonitorService) UpdateMonitor(monitorHigh, monitorLow float64, id int64, uid uint) (err error) {
	err = dao.StockMonitorDao.UpdateByMonitorNum(monitorHigh, monitorLow, id, uid)
	if err != nil {
		return
	}
	return
}
