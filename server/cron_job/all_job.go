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
		for _, stock := range stockList {
			stock.CreateTime = time.Now().Unix()
			_ = dao.StockDao.Save(stock)
		}
	}
}
