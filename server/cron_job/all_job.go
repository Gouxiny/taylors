package cron_job

import (
	"taylors/dao"
	"taylors/model"
	"taylors/service"
	"time"
)

var AllJob *allJob

type allJob struct {
	List []model.Stock
}

func init() {
	AllJob = new(allJob)
}

func (job *allJob) Run() {
	defer func() {
		recover()
	}()

	stockList, err := service.StockAllService.AllDetailList(1, 100000)
	if err != nil {
		return
	}

	if stockList != nil && len(stockList) > 0 {
		maxId, err := dao.StockDao.Max()
		if err != nil {
			return
		}

		flag := false
		if maxId < 1 {
			//获取历史数据
			flag = true
		}

		for _, stock := range stockList {
			if flag {
				historyList, err := service.StockCommonService.HistoryList(stock.Code)
				if err != nil {
					continue
				}

				for _, historyStock := range historyList {
					historyStock.Code = stock.Code
					historyStock.Name = stock.Name
					historyStock.Symbol = stock.Symbol
					historyStock.Exchange = stock.Exchange
					historyStock.TotalShares = stock.TotalShares
					historyStock.CreateTime = time.Now().Unix()

					err := dao.StockDao.Save(historyStock)
					if err != nil {
						continue
					}
				}
			}
			stock.CreateTime = time.Now().Unix()
			err := dao.StockDao.Save(stock)
			if err != nil {
				continue
			}
		}
	}
}
