package service

import (
	"taylors/model"
	"taylors_proto/taylors_stock"
	"time"
)

var StockTopService = new(stockTopService)
var StockAllService = new(stockAllService)
var StockMonitorService = new(stockMonitorService)
var StockCommonService = new(stockCommonService)

var _OverTime = time.Second * 10

func Conv(gstocks []*taylors_stock.Stock) (stockList []model.Stock) {
	for _, top := range gstocks {
		stockList = append(stockList, model.Stock{
			Symbol:             top.Symbol + "-" + top.Name,
			Name:               top.Name,
			Exchange:           top.Exchange,
			Code:               top.Code,
			TotalShares:        top.TotalShares,
			MarketCapital:      top.MarketCapital,
			Pb:                 top.Pb,
			PeTtm:              top.PeTtm,
			CurrentYearPercent: top.CurrentYearPercent,
			High52w:            top.High52W,
			Low52w:             top.Low52W,
			LimitDown:          top.LimitDown,
			High:               top.High,
			Chg:                top.Chg,
			Low:                top.Low,
			Open:               top.Open,
			LastClose:          top.LastClose,
			Volume:             int(top.Volume),
			Amount:             top.Amount,
			Percent:            top.Percent,
			VolumeRatio:        top.VolumeRatio,
			Time:               top.Time,
		})
	}
	return
}
