package service

import (
	"taylors/crawler"
	"taylors/dao"
	"taylors/model"
	"taylors/model/param"
	"time"
)

type stockAllService struct {
}

func (*stockAllService) AllList(allListParam *param.AllListParam) (stockList []*model.Stock, total int, err error) {
	maxStock, err := dao.StockDao.Max()
	if err != nil {
		return
	}
	timeStock := time.Unix(maxStock.CreateTime, 0)
	t := timeStock.Add(-time.Hour * 3)

	allListParam.CreateTime = t.Unix()
	stockList, total, err = dao.StockDao.AllList(allListParam)
	return
}

func (*stockAllService) AllListByCrawler() (stockList []model.Stock, err error) {
	stockList = crawler.NewDongFangCrawler().All()
	return
}
