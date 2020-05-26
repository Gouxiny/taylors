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

	if service.StockCommonService.CheckOffday() { // 判断是否是休市日
		return
	}

	stockList, err := service.StockAllService.AllDetailList()
	if err != nil {
		return
	}

	if stockList != nil && len(stockList) > 0 {
		maxId, err := dao.StockDao.Max()
		if err != nil {
			return
		}

		flag := false
		//if maxId < 1 {
		//	//获取历史数据
		//	flag = true
		//}
		_ = maxId

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
