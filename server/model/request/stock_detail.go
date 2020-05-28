package request

type StockTopListReq struct {
	MarketCapitalMax float64 `json:"marketCapitalMax" desc:"市值范围最大" `
	MarketCapitalMin float64 `json:"marketCapitalMin" desc:"市值范围最小" `
	PercentMax       float64 `json:"percentMax" desc:"涨幅范围最大" `
	PercentMin       float64 `json:"percentMin" desc:"涨幅范围最小" `
	VolumeRatioMax   float64 `json:"volume_ratio_max" desc:"量比范围最大" `
	VolumeRatioMin   float64 `json:"volume_ratio_min" desc:"量比范围最小" `
	CurrentMax       float64 `json:"currentMax" desc:"股价范围最大" `
	CurrentMin       float64 `json:"currentMin" desc:"股价范围最小" `
}

type AllDetailListReq struct {
	PageNum  int `json:"page_num" desc:"页数" `
	PageSize int `json:"page_size" desc:"每页数量" `
}

type MonitorOneReq struct {
	Id int64 `json:"id" desc:"key" `
}

type MonitorListReq struct {
	Name             string  `json:"name" desc:"名称" `
	Symbol           string  `json:"symbol" desc:"编码" `
	MarketCapitalMax float64 `json:"marketCapitalMax" desc:"市值范围最大" `
	MarketCapitalMin float64 `json:"marketCapitalMin" desc:"市值范围最小" `
	PercentMax       float64 `json:"percentMax" desc:"涨幅范围最大" `
	PercentMin       float64 `json:"percentMin" desc:"涨幅范围最小" `
	VolumeRatioMax   float64 `json:"volume_ratio_max" desc:"量比范围最大" `
	VolumeRatioMin   float64 `json:"volume_ratio_min" desc:"量比范围最小" `
	CurrentMax       float64 `json:"currentMax" desc:"股价范围最大" `
	CurrentMin       float64 `json:"currentMin" desc:"股价范围最小" `
}

type AddMonitorReq struct {
	Symbol      string  `json:"symbol" desc:"编码" `
	MonitorHigh float64 `json:"monitor_high" desc:"高位预警" `
	MonitorLow  float64 `json:"monitor_low" desc:"低位预警" `
}

type DelMonitorReq struct {
	Id int64 `json:"id" desc:"key" `
}

type UpdateMonitorReq struct {
	Id          int64   `json:"id" desc:"key" `
	MonitorHigh float64 `json:"monitor_high" desc:"高位预警" `
	MonitorLow  float64 `json:"monitor_low" desc:"低位预警" `
}
