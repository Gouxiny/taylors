package dao

var StockDao *newStock

func Init() {
	StockDao = newStockDao()
}
