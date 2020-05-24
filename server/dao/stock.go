package dao

import (
	"github.com/jinzhu/gorm"
	"taylors/global"
	"taylors/model"
)

type newStock struct {
	Db *gorm.DB
}

func newStockDao() *newStock {
	stock := &model.Stock{}
	return &newStock{Db: global.GVA_DB.Model(stock)}
}

func (dao *newStock) Save(stock model.Stock) (err error) {
	err = dao.Db.Save(&stock).Error
	return
}

func (dao *newStock) Max() (maxId int64, err error) {
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
