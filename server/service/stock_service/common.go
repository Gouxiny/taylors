package stock_service

import (
	"gin-vue-admin/model"
	"taylors_proto/taylors_stock"
)

var StockService *stockService

func init() {
	StockService = new(stockService)
}

func Conv(gstocks []*taylors_stock.Stock) (stockList []model.Stock) {
	for _, top := range gstocks {
		stockList = append(stockList, model.Stock{
			Symbol:             top.Symbol,
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
			Time:               top.Time,
		})
	}
	return
}
