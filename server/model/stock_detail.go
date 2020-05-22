// 自动生成模板Stock
package model

import (
	"github.com/jinzhu/gorm"
)

type Stock struct {
	gorm.Model
	Id                 int64   `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Symbol             string  `gorm:"column:symbol"`               //编码
	Name               string  `gorm:"column:name"`                 //名称
	Exchange           string  `gorm:"column:exchange"`             //上市地
	Code               string  `gorm:"column:code"`                 //股票代码
	TotalShares        int64   `gorm:"column:total_shares"`         //总股本
	MarketCapital      float64 `gorm:"column:market_capital"`       //总市值
	Pb                 float64 `gorm:"column:pb"`                   //市净率
	PeTtm              float64 `gorm:"column:pe_ttm"`               //市盈率(TTM)
	CurrentYearPercent float64 `gorm:"column:current_year_percent"` //今年涨幅
	High52w            float64 `gorm:"column:high52w"`              //52周最高
	Low52w             float64 `gorm:"column:low52w"`               //52周最低
	LimitDown          float64 `gorm:"column:limit_down"`           //跌停
	High               float64 `gorm:"column:high"`                 //涨停
	Chg                float64 `gorm:"column:chg"`                  //最高(当天)
	Low                float64 `gorm:"column:low"`                  //最低(当天)
	Open               float64 `gorm:"column:open"`                 //今开
	LastClose          float64 `gorm:"column:last_close"`           //昨收
	Volume             int     `gorm:"column:volume"`               //成交量
	Amount             float64 `gorm:"column:amount"`               //成交额
	Percent            float64 `gorm:"column:percent"`              //实时涨幅
	Time               int64   `gorm:"column:time"`                 //时间
	CreateTime         int64   `gorm:"column:createTime"`           //数据创建时间
}
