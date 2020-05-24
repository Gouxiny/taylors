package request

type StockTopListReq struct {
	MarketCapitalMax float64 `json:"marketCapitalMax" desc:"市值范围最大" `
	MarketCapitalMin float64 `json:"marketCapitalMin" desc:"市值范围最小" `
	PercentMax       float64 `json:"percentMax" desc:"涨幅范围最大" `
	PercentMin       float64 `json:"percentMin" desc:"涨幅范围最小" `
	VolumeRatioMax   float64 `json:"volume_ratio_max" desc:"量比范围最大" `
	VolumeRatioMin   float64 `json:"volume_ratio_min" desc:"量比范围最小" `
}

type MonitorListReq struct {
	SymbolList []string `json:"symbol_list" desc:"编码" `
}

type AllDetailListReq struct {
	PageNum  int `json:"page_num" desc:"页数" `
	PageSize int `json:"page_size" desc:"每页数量" `
}
