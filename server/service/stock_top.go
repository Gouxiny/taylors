package service

import (
	"context"
	"taylors/model"
	"taylors/model/request"
	"taylors/module/crawler"
	"taylors_proto/taylors_stock"
)

type stockTopService struct {
}

func (*stockTopService) TopList(filter request.StockTopListReq) (stockList []model.Stock, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), _OverTime)
	defer cancel()
	req := &taylors_stock.TopReq{}
	topRsp, err := crawler.Grpc_cli.ListTop(ctx, req)
	if err != nil {
		return
	}
	stocks := Conv(topRsp.StockList)

	for _, stock := range stocks {
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
		stockList = append(stockList, stock)
	}

	return
}
