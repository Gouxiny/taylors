package dao

import (
	"github.com/jinzhu/gorm"
	"taylors/global"
	"taylors/model"
)

type stockModel struct {
	Db *gorm.DB
}

func newStockDao() *stockModel {
	stock := &model.Stock{}
	global.GVA_DB.Model(&model.Stock{}).AddIndex("base_index", "symbol", "name", "exchange", "code")
	return &stockModel{Db: global.GVA_DB.Model(stock)}
}

func (dao *stockModel) Save(stock model.Stock) (err error) {
	err = dao.Db.Save(&stock).Error
	return
}

func (dao *stockModel) FindBySymbol(symbol string) (stockList []*model.Stock, err error) {
	err = dao.Db.Where("symbol = ? ", symbol).Order("create_time desc").Find(&stockList).Error
	return
}

func (dao *stockModel) Max() (maxId int64, err error) {
	type C struct {
		MaxId int64 `gorm:"column:max_id"`
	}
	count := &C{}
	err = dao.Db.Raw("SELECT max(id) AS max_id from stock").Scan(count).Error
	if err != nil {
		return
	}
	maxId = count.MaxId
	return
}
