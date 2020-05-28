package service

import (
	"taylors/global"
	"taylors/model"
	"taylors_proto/taylors_stock"
	"time"
)

var StockTopService = new(stockTopService)
var StockAllService = new(stockAllService)
var StockMonitorService = new(stockMonitorService)
var StockCommonService = new(stockCommonService)

var _OverTime = time.Second * 10

func Init() {
	err := StockMonitorService.SyncMonitor()
	if err != nil {
		global.GVA_LOG.Error("初始化监控失败", err)
	}
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
			Current:            top.Current,
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

func ConvStockMonitorModel(stock model.Stock) (stockModel model.StockMonitorModel) {
	stockModel = model.StockMonitorModel{
		Id:                 stock.Id,
		Symbol:             stock.Symbol,
		Name:               stock.Name,
		Exchange:           stock.Exchange,
		Code:               stock.Code,
		TotalShares:        stock.TotalShares,
		MarketCapital:      stock.MarketCapital,
		Current:            stock.Current,
		Pb:                 stock.Pb,
		PeTtm:              stock.PeTtm,
		CurrentYearPercent: stock.CurrentYearPercent,
		High52w:            stock.High52w,
		Low52w:             stock.Low52w,
		LimitDown:          stock.LimitDown,
		High:               stock.High,
		Chg:                stock.Chg,
		Low:                stock.Low,
		Open:               stock.Open,
		LastClose:          stock.LastClose,
		Volume:             stock.Volume,
		Amount:             stock.Amount,
		Percent:            stock.Percent,
		VolumeRatio:        stock.VolumeRatio,
		Time:               stock.Time,
	}
	return
}
