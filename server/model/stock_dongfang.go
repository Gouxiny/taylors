// 自动生成模板Stock
package model

type Stock struct {
	Id int64 `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	F1 int   `gorm:"column:f1" json:"f1"`

	Current      float64 `gorm:"column:current" json:"f2"`       //最新价
	Percent      float64 `gorm:"column:percent" json:"f3"`       //涨跌幅
	Chg          float64 `gorm:"column:chg" json:"f4"`           //涨跌额
	Volume       int64   `gorm:"column:volume" json:"f5"`        //成交量（万手）
	Amount       float64 `gorm:"column:amount" json:"f6"`        //成交额
	Amplitude    float64 `gorm:"column:amplitude" json:"f7"`     //振幅
	TurnoverRate float64 `gorm:"column:turnover_rate" json:"f8"` //换手率
	PeTtm        float64 `gorm:"column:pe_ttm" json:"f9"`        //市盈率
	VolumeRatio  float64 `gorm:"column:volume_ratio" json:"f10"` //量比

	F11 int `gorm:"column:f11" json:"f11"`

	Code              string  `gorm:"column:code" json:"f12"`               //编码
	Szh               int     `gorm:"column:szh" json:"f13"`                //0:sz 1:sh
	Name              string  `gorm:"column:name" json:"f14"`               //名字
	High              float64 `gorm:"column:high" json:"f15"`               //最高
	Low               float64 `gorm:"column:low" json:"f16"`                //最低
	Open              float64 `gorm:"column:open" json:"f17"`               //今开
	LastClose         float64 `gorm:"column:last_close" json:"f18"`         //昨收
	MarketCapital     int64   `gorm:"column:market_capital" json:"f20"`     //总市值
	CirculationMarket int64   `gorm:"column:circulation_market" json:"f21"` //流通市值

	F22 int `gorm:"column:f22" json:"f22"`

	Pb          float64 `gorm:"column:pb" json:"f23"`           //市净率
	Percent60   float64 `gorm:"column:percent_60" json:"f24"`   //60日涨幅
	PercentYear float64 `gorm:"column:percent_year" json:"f25"` //年初至今涨跌幅

	F62  int64   `gorm:"column:f62" json:"f62"`
	F115 float64 `gorm:"column:f115" json:"f115"`

	F128 string `gorm:"column:f128" json:"f128"`
	F140 string `gorm:"column:f140" json:"f140"`
	F141 string `gorm:"column:f141" json:"f141"`
	F136 string `gorm:"column:f136" json:"f136"`

	F152       int   `gorm:"column:f152" json:"f152"`
	CreateTime int64 `gorm:"column:create_time" json:"create_time"`
}

func (Stock) TableName() string {
	return "stock"
}
