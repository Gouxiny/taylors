package service

import (
	"sort"
	"sync"
	"taylors/dao"
	"taylors/model"
	"taylors/model/param"
)

type analysisModel struct {
	score float32
	code  string
}

type analysisModelList []*analysisModel

func (a analysisModelList) Len() int {
	return len(a)
}
func (a analysisModelList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a analysisModelList) Less(i, j int) bool { return a[i].score < a[j].score }

type stockAnalysisService struct {
}

func (analy *stockAnalysisService) AnalysisList(filter *param.AnalysisListParam) (stockList []*model.Stock, total int, err error) {
	stockCodeList, err := dao.StockDao.CodeList()
	if err != nil {
		return
	}

	scoreList := make([]*analysisModel, 0, 5000)
	analysisCh := make(chan *analysisModel, 100)

	go func() {
		for analysis := range analysisCh {
			if analysis != nil && analysis.code != "" {
				scoreList = append(scoreList, analysis)
			}
		}
	}()

	wg := sync.WaitGroup{}

	for _, stock := range stockCodeList {
		wg.Add(1)
		go func(stockObj *model.Stock) {
			defer wg.Done()
			stockPOList, err := dao.StockDao.FindByCode(stockObj.Code)
			if err != nil {
				return
			}
			calculateScore := analy.CalculateScore(stockPOList)
			analysisCh <- &analysisModel{
				score: calculateScore,
				code:  stockObj.Code,
			}
		}(stock)
	}
	wg.Wait()

	sort.Sort(analysisModelList(scoreList))

	return
}

func (*stockAnalysisService) CalculateScore(stockList []*model.Stock) (score float32) {
	if stockList == nil || len(stockList) == 0 {
		return
	}

	for _, stock := range stockList {
		score += float32(stock.Current)
	}

	return
}
