package service

import (
	"fmt"
	"taylors/dao"
	"taylors/model"
	"taylors/model/param"
)

type stockAnalysisService struct {
}

func (*stockAnalysisService) AnalysisList(filter *param.AnalysisListParam) (stockList []*model.Stock, total int, err error) {
	stockCodeList, err := dao.StockDao.CodeList()
	if err != nil {
		return
	}
	for _, stock := range stockCodeList {
		stockPOList, err := dao.StockDao.FindByCode(stock.Code)
		if err != nil {
			break
		}
		for _, stockPO := range stockPOList {
			fmt.Println(stockPO)
		}

	}
	return
}
