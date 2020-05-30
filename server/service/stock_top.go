package service

import (
	"taylors/crawler"
	"taylors/model"
	"taylors/model/request"
)

type stockTopService struct {
}

func (*stockTopService) TopList(filter request.StockTopListReq) (stockList []model.Stock, err error) {
	stockListCrawler := crawler.NewDongFangCrawler().Top()

	for _, stock := range stockListCrawler {
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
		stock.MarketCapital = stock.MarketCapital / 100000000
		stockList = append(stockList, stock)
	}

	return
}
