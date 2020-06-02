package service

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"sort"
	"sync"
	"taylors/dao"
	"taylors/model"
	"taylors/model/param"
)

type stockAnalysisService struct {
}

func (analy *stockAnalysisService) AnalysisList(filter *param.AnalysisListParam) (stockList []*model.Stock, total int, err error) {
	stockCodeList, err := dao.StockDao.CodeList()
	if err != nil {
		return
	}

	page := filter.Page
	if page < 1 {
		page = 1
	}
	pageSize := filter.PageSize
	if pageSize < 1 {
		pageSize = 1
	}
	filter.Page = 0
	filter.PageSize = 0
	filterBys, err := json.Marshal(filter)
	if err != nil {
		return
	}
	hash := md5.New()
	hash.Write(filterBys)
	md5 := hex.EncodeToString(hash.Sum(nil))

	exitCache := false
	analysisModelCacheList := NewAnalysisCache().Get(md5)

	if len(analysisModelCacheList) > 0 {
		exitCache = true
	}

	if !exitCache {
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

		NewAnalysisCache().Cache(md5, scoreList)
		analysisModelCacheList = scoreList
	}

	endIndex := page * pageSize
	startIndex := (page - 1) * pageSize
	if endIndex > len(analysisModelCacheList)-1 {
		endIndex = len(analysisModelCacheList) - 1
	}

	codes := []string{}
	for i := startIndex; i < endIndex; i++ {
		codes = append(codes, analysisModelCacheList[i].code)
	}

	total = len(analysisModelCacheList)

	for _, code := range codes {
		stockPO, err := dao.StockDao.FindLastByCode(code)
		if err != nil {
			stockList = nil
			break
		}
		stockList = append(stockList, stockPO)
	}
	return
}

func (*stockAnalysisService) CalculateScore(stockList []*model.Stock) (score float32) {
	if stockList == nil || len(stockList) == 0 {
		return
	}

	for _, stock := range stockList {
		score += float32(stock.Percent)
	}

	return
}
