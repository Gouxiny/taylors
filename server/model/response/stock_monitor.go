package response

type StockMonitorModel struct {
	Id                 int64   `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Symbol             string  `gorm:"column:symbol" json:"symbol"`                             //编码
	Name               string  `gorm:"column:name" json:"name"`                                 //名称
	Exchange           string  `gorm:"column:exchange" json:"exchange"`                         //上市地
	Code               string  `gorm:"column:code" json:"code"`                                 //股票代码
	TotalShares        int64   `gorm:"column:total_shares" json:"total_shares"`                 //总股本
	MarketCapital      float64 `gorm:"column:market_capital" json:"market_capital"`             //总市值
	Current            float64 `gorm:"column:current" json:"current"`                           //当前价
	Pb                 float64 `gorm:"column:pb" json:"pb"`                                     //市净率
	PeTtm              float64 `gorm:"column:pe_ttm" json:"pe_ttm"`                             //市盈率(TTM)
	CurrentYearPercent float64 `gorm:"column:current_year_percent" json:"current_year_percent"` //今年涨幅
	High52w            float64 `gorm:"column:high52w" json:"high_52_w"`                         //52周最高
	Low52w             float64 `gorm:"column:low52w" json:"low_52_w"`                           //52周最低
	LimitDown          float64 `gorm:"column:limit_down" json:"limit_down"`                     //跌停
	High               float64 `gorm:"column:high" json:"high"`                                 //涨停
	Chg                float64 `gorm:"column:chg" json:"chg"`                                   //最高(当天)
	Low                float64 `gorm:"column:low" json:"low"`                                   //最低(当天)
	Open               float64 `gorm:"column:open" json:"open"`                                 //今开
	LastClose          float64 `gorm:"column:last_close" json:"last_close"`                     //昨收
	Volume             int     `gorm:"column:volume" json:"volume"`                             //成交量
	Amount             float64 `gorm:"column:amount" json:"amount"`                             //成交额
	Percent            float64 `gorm:"column:percent" json:"percent"`                           //实时涨幅
	VolumeRatio        float64 `gorm:"volume_ratio" json:"volume_ratio"`                        //量比
	Time               int64   `gorm:"column:time" json:"time"`                                 //时间
	CreateTime         int64   `gorm:"column:create_time" json:"create_time"`                   //数据创建时间
	MonitorHigh        float64 `gorm:"column:monitor_high" json:"monitor_high"`                 //高位预警
	MonitorLow         float64 `gorm:"column:monitor_low" json:"monitor_low"`                   //低位预警
}
