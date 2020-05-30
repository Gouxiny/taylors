package service

import (
	"taylors/crawler"
	"taylors/dao"
)

type stockCommonService struct {
}

func (srv *stockCommonService) CheckOffday() (offday bool) {
	codes := []string{"000001"}
	stockList := crawler.NewDongFangCrawler().Monitor(codes)

	if len(stockList) < 1 {
		return
	}

	stock := stockList[0]

	stockPOList, err := dao.StockDao.FindByCode(stock.Code)
	if err != nil {
		return
	}
	if len(stockPOList) == 0 {
		return
	}

	stockPO := stockPOList[0]

	if stock.Code != stockPO.Code {
		return
	}

	if stock.Percent != stockPO.Percent {
		return
	}

	if stock.Amount != stockPO.Amount {
		return
	}

	if stock.Volume != stockPO.Volume {
		return
	}

	offday = true

	return
}
