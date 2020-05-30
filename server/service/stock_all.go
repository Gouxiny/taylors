package service

import (
	"taylors/crawler"
	"taylors/model"
)

type stockAllService struct {
}

func (*stockAllService) AllList() (stockList []model.Stock, err error) {
	stockList = crawler.NewDongFangCrawler().All()
	return
}
