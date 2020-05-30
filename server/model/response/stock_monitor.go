package response

import "taylors/model"

type StockMonitorModel struct {
	model.Stock
	MonitorHigh float64 `gorm:"column:monitor_high" json:"monitor_high"` //高位预警
	MonitorLow  float64 `gorm:"column:monitor_low" json:"monitor_low"`   //低位预警
}
