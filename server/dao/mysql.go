package dao

var StockDao *newStock

func init() {
	StockDao = newStockDao()
}
