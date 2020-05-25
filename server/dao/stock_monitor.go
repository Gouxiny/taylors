package dao

import (
	"github.com/jinzhu/gorm"
	"taylors/global"
	"taylors/model"
	"time"
)

type stockMonitorModel struct {
	Db *gorm.DB
}

func newStockMonitorDao() *stockMonitorModel {
	monitor := &model.StockMonitor{}
	global.GVA_DB.Model(&model.StockMonitor{}).AddUniqueIndex("base_index", "symbol", "del_status", "user_id")
	return &stockMonitorModel{Db: global.GVA_DB.Model(monitor)}
}

func (dao *stockMonitorModel) Save(stock model.StockMonitor) (err error) {
	err = dao.Db.Save(&stock).Error
	return
}

func (dao *stockMonitorModel) DelBySymbol(symbol string, userId uint) (err error) {
	err = dao.Db.Exec("update stock_monitor set del_status = ? where symbol = ? and user_id = ?", time.Now().Unix(), symbol, userId).Error
	return
}

func (dao *stockMonitorModel) UpdateByMonitorNum(monitorHigh, monitorLow float64, id int64, uid uint) (err error) {
	err = dao.Db.Exec("update stock_monitor set monitor_high = ?,set monitor_low = ? where id = ? and user_id = ? ", monitorHigh, monitorLow, id, uid).Error
	return
}

func (dao *stockMonitorModel) ListByUser(userId uint) (stockMonitorList []*model.StockMonitor, err error) {
	err = dao.Db.Where("user_id = ? AND del_status = ? ", userId, DEL_STATUS).Order("create_time desc").Find(&stockMonitorList).Error
	return
}

func (dao *stockMonitorModel) FindByUserAndSymbol(userId uint, symbol string) (stockMonitor *model.StockMonitor, err error) {
	err = dao.Db.Where("user_id = ? and symbol = ? AND del_status = ? ", userId, symbol, DEL_STATUS).Order("create_time desc").First(&stockMonitor).Error
	return
}

func (dao *stockMonitorModel) ListNotDel() (stockMonitorList []*model.StockMonitor, err error) {
	err = dao.Db.Where("del_status = ? ", DEL_STATUS).Order("create_time desc").Find(&stockMonitorList).Error
	return
}
