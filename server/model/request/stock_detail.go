package request

type StockTopListReq struct {
	DataSize     int     `json:"data_size" desc:"数量" `
	CapitalStart float64 `json:"capital_start" desc:"市值范围开始" `
	CapitalEnd   float64 `json:"capital_end" desc:"市值范围结束" `
}

type MonitorListReq struct {
	SymbolList []string `json:"symbol_list" desc:"编码" `
}
