package dao

var StockDao *stockModel
var StockMonitorDao *stockMonitorModel

func Init() {
	StockDao = newStockDao()
	StockMonitorDao = newStockMonitorDao()
}
