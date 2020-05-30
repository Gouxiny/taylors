package crawler

import (
	"sync"
	"taylors/model"
)

type dongFangCache struct {
	lock        sync.RWMutex
	dongFangMap map[string]model.Stock
}

var cacheOnce sync.Once
var dongFangCacheObj *dongFangCache

func instence() *dongFangCache {
	cacheOnce.Do(func() {
		dongFangCacheObj = &dongFangCache{
			dongFangMap: make(map[string]model.Stock, 5000),
		}
	})
	return dongFangCacheObj
}

func (cache *dongFangCache) store(dongFang *dongFang) {
	cache.lock.Lock()
	defer cache.lock.Unlock()
	for _, stock := range dongFang.Data.Diff {
		cache.dongFangMap[stock.Code] = stock
	}
}

func (cache *dongFangCache) obtain(codes []string) (stockList []model.Stock) {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	for _, code := range codes {
		model, ok := cache.dongFangMap[code]
		if ok {
			stockList = append(stockList, model)
		}
	}
	return
}

func (cache *dongFangCache) obtainAll() (stockList []model.Stock) {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	for key, dongFangModel := range cache.dongFangMap {
		if key != "" {
			stockList = append(stockList, dongFangModel)
		}
	}

	return
}
