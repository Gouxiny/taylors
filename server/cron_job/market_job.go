package cron_job

import (
	"taylors/dao"
	"taylors/model"
	"taylors/service"
	"time"
)

var MarketJob *marketJob

type marketJob struct {
	List []model.Stock
}

func init() {
	MarketJob = new(marketJob)
}

func (job *marketJob) Run() {
	defer func() {
		recover()
	}()

	if service.StockCommonService.CheckOffday() { // 判断是否是休市日
		return
	}

	stockList, err := service.StockAllService.AllList()
	if err != nil {
		return
	}

	if stockList != nil && len(stockList) > 0 {
		for _, stock := range stockList {
			stock.CreateTime = time.Now().Unix()
			_ = dao.StockDao.Save(stock)
		}
	}
}
