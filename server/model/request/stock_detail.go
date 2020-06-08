package request

type StockTopListReq struct {
	Name             string  `json:"name" desc:"名称" `
	Code             string  `json:"code" desc:"编码" `
	MarketCapitalMax float64 `json:"marketCapitalMax" desc:"市值范围最大" `
	MarketCapitalMin float64 `json:"marketCapitalMin" desc:"市值范围最小" `
	PercentMax       float64 `json:"percentMax" desc:"涨幅范围最大" `
	PercentMin       float64 `json:"percentMin" desc:"涨幅范围最小" `
	VolumeRatioMax   float64 `json:"volume_ratio_max" desc:"量比范围最大" `
	VolumeRatioMin   float64 `json:"volume_ratio_min" desc:"量比范围最小" `
	CurrentMax       float64 `json:"currentMax" desc:"股价范围最大" `
	CurrentMin       float64 `json:"currentMin" desc:"股价范围最小" `
}

type AllListReq struct {
	PageInfo
	Name             string  `json:"name" desc:"名称" `
	Code             string  `json:"code" desc:"编码" `
	MarketCapitalMax float64 `json:"marketCapitalMax" desc:"市值范围最大" `
	MarketCapitalMin float64 `json:"marketCapitalMin" desc:"市值范围最小" `
	PercentMax       float64 `json:"percentMax" desc:"涨幅范围最大" `
	PercentMin       float64 `json:"percentMin" desc:"涨幅范围最小" `
	VolumeRatioMax   float64 `json:"volume_ratio_max" desc:"量比范围最大" `
	VolumeRatioMin   float64 `json:"volume_ratio_min" desc:"量比范围最小" `
	CurrentMax       float64 `json:"currentMax" desc:"股价范围最大" `
	CurrentMin       float64 `json:"currentMin" desc:"股价范围最小" `
}

type MonitorOneReq struct {
	Id int64 `json:"id" desc:"key" `
}

type MonitorListReq struct {
	IsDay            bool    `json:"isDay" desc:"是否日监控" `
	Name             string  `json:"name" desc:"名称" `
	Code             string  `json:"code" desc:"编码" `
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
	IsDay       bool    `json:"isDay" desc:"是否日监控" `
	Code        string  `json:"code" desc:"编码" `
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

type AnalysisListReq struct {
	PageInfo
	Name             string  `json:"name" desc:"名称" `
	Code             string  `json:"code" desc:"编码" `
	MarketCapitalMax float64 `json:"marketCapitalMax" desc:"市值范围最大" `
	MarketCapitalMin float64 `json:"marketCapitalMin" desc:"市值范围最小" `
	PercentMax       float64 `json:"percentMax" desc:"涨幅范围最大" `
	PercentMin       float64 `json:"percentMin" desc:"涨幅范围最小" `
	VolumeRatioMax   float64 `json:"volume_ratio_max" desc:"量比范围最大" `
	VolumeRatioMin   float64 `json:"volume_ratio_min" desc:"量比范围最小" `
	CurrentMax       float64 `json:"currentMax" desc:"股价范围最大" `
	CurrentMin       float64 `json:"currentMin" desc:"股价范围最小" `
	StartTime        int64   `json:"startTime" desc:"开始时间" `
	EndTime          int64   `json:"endTime" desc:"结束时间" `
	DayMin           int     `json:"dayMin" desc:"最少天数" `
}
