package service

import (
	"taylors/dao"
	"taylors/model"
	"taylors/model/param"
)

type stockAnalysisService struct {
}

func (*stockAnalysisService) AnalysisList(filter *param.AnalysisListParam) (stockList []*model.Stock, total int, err error) {
	stockList, err = dao.StockDao.AnalysisList(filter)
	if err != nil {
		return
	}

	return
}
