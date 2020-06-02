package dao

import (
	"github.com/jinzhu/gorm"
	"taylors/global"
	"taylors/model"
	"taylors/model/param"
)

type stockModel struct {
	Db *gorm.DB
}

func newStockDao() *stockModel {
	stock := &model.Stock{}
	global.GVA_DB.Model(&model.Stock{}).AddIndex("base_index", "code", "name", "percent", "szh", "current")
	return &stockModel{Db: global.GVA_DB.Model(stock)}
}

func (dao *stockModel) Save(stock model.Stock) (err error) {
	err = dao.Db.Save(&stock).Error
	return
}

func (dao *stockModel) FindByCode(code string) (stockList []*model.Stock, err error) {
	err = dao.Db.Where("code = ? ", code).Order("create_time desc").Find(&stockList).Error
	return
}

func (dao *stockModel) FindLastByCode(code string) (stock *model.Stock, err error) {
	stock = &model.Stock{}
	err = dao.Db.Where("code = ? ", code).Order("create_time desc").First(stock).Error
	return
}

func (dao *stockModel) Max() (stock *model.Stock, err error) {
	stock = &model.Stock{}
	err = dao.Db.Order("create_time desc").Limit(1).First(stock).Error
	return
}

func (dao *stockModel) AllList(allListParam *param.AllListParam) (stockList []*model.Stock, total int, err error) {
	limit := allListParam.PageSize
	offset := allListParam.PageSize * (allListParam.Page - 1)
	db := dao.Db

	if allListParam.Code != "" {
		db = db.Where("code LIKE ?", "%"+allListParam.Code+"%")
	}
	if allListParam.Name != "" {
		db = db.Where("name LIKE ?", "%"+allListParam.Name+"%")
	}

	if allListParam.MarketCapitalMax != 0 {
		db = db.Where("market_capital <= ?", allListParam.MarketCapitalMax)
	}
	if allListParam.MarketCapitalMin != 0 {
		db = db.Where("market_capital >= ?", allListParam.MarketCapitalMin)
	}

	if allListParam.PercentMax != 0 {
		db = db.Where("percent <= ?", allListParam.PercentMax)
	}
	if allListParam.PercentMin != 0 {
		db = db.Where("percent >= ?", allListParam.PercentMin)
	}

	if allListParam.VolumeRatioMax != 0 {
		db = db.Where("volume_ratio <= ?", allListParam.VolumeRatioMax)
	}
	if allListParam.VolumeRatioMin != 0 {
		db = db.Where("volume_ratio >= ?", allListParam.VolumeRatioMin)
	}

	if allListParam.CurrentMax != 0 {
		db = db.Where("current <= ?", allListParam.CurrentMax)
	}
	if allListParam.CurrentMin != 0 {
		db = db.Where("current >= ?", allListParam.CurrentMin)
	}

	if allListParam.CreateTime != 0 {
		db = db.Where("create_time >= ?", allListParam.CreateTime)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	db = db.Limit(limit).Offset(offset)
	err = db.Order("code", true).Find(&stockList).Error
	return
}

func (dao *stockModel) CodeList() (stockList []*model.Stock, err error) {
	err = dao.Db.Raw("select code from stock where batch_code = (select batch_code from stock order by id DESC limit 1)").Scan(&stockList).Error
	return
}

func (dao *stockModel) AnalysisList(filter *param.AnalysisListParam) (stockList []*model.Stock, err error) {
	db := dao.Db

	if filter.Code != "" {
		db = db.Where("code LIKE ?", "%"+filter.Code+"%")
	}
	if filter.Name != "" {
		db = db.Where("name LIKE ?", "%"+filter.Name+"%")
	}

	if filter.MarketCapitalMax != 0 {
		db = db.Where("market_capital <= ?", filter.MarketCapitalMax)
	}
	if filter.MarketCapitalMin != 0 {
		db = db.Where("market_capital >= ?", filter.MarketCapitalMin)
	}

	if filter.PercentMax != 0 {
		db = db.Where("percent <= ?", filter.PercentMax)
	}
	if filter.PercentMin != 0 {
		db = db.Where("percent >= ?", filter.PercentMin)
	}

	if filter.VolumeRatioMax != 0 {
		db = db.Where("volume_ratio <= ?", filter.VolumeRatioMax)
	}
	if filter.VolumeRatioMin != 0 {
		db = db.Where("volume_ratio >= ?", filter.VolumeRatioMin)
	}

	if filter.CurrentMax != 0 {
		db = db.Where("current <= ?", filter.CurrentMax)
	}
	if filter.CurrentMin != 0 {
		db = db.Where("current >= ?", filter.CurrentMin)
	}

	err = db.Order("code", true).Find(&stockList).Error
	return
}
