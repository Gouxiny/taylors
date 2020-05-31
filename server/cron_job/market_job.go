package cron_job

import (
	"github.com/satori/go.uuid"
	"taylors/dao"
	"taylors/service"
	"time"
)

var MarketJob *marketJob

type marketJob struct {
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

	stockList, err := service.StockAllService.AllListByCrawler()
	if err != nil {
		return
	}

	batchCode := uuid.NewV4().String()
	if stockList != nil && len(stockList) > 0 {
		for _, stock := range stockList {
			stock.CreateTime = time.Now().Unix()
			stock.BatchCode = batchCode
			_ = dao.StockDao.Save(stock)
		}
	}

}
