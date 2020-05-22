package dao

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"github.com/jinzhu/gorm"
)

type newStock struct {
	Db *gorm.DB
}

func newStockDao() *newStock {
	stock := &model.Stock{}
	if err := global.GVA_DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(stock).Error; err != nil {
		panic(err)
	}

	global.GVA_DB.Model(stock).AddIndex("base_index", "symbol", "name", "exchange", "code")

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
