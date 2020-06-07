package service

import "sync"

type analysisModel struct {
	score float32
	code  string
}

type analysisModelList []*analysisModel

func (a analysisModelList) Len() int {
	return len(a)
}
func (a analysisModelList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a analysisModelList) Less(i, j int) bool { return a[i].score > a[j].score }

type analysisCache struct {
	data map[string][]*analysisModel
	lock sync.Mutex
}

var AnalysisCache *analysisCache
var analysisCacheOnce sync.Once

func NewAnalysisCache() *analysisCache {
	analysisCacheOnce.Do(func() {
		AnalysisCache = &analysisCache{
			data: make(map[string][]*analysisModel),
		}
	})
	return AnalysisCache
}

func (analy *analysisCache) Cache(code string, analysisModelList []*analysisModel) {
	analy.lock.Lock()
	defer analy.lock.Unlock()
	analy.data[code] = analysisModelList
}

func (analy *analysisCache) Get(code string) (analysisModelList []*analysisModel) {
	analy.lock.Lock()
	defer analy.lock.Unlock()
	cacheData, ok := analy.data[code]
	if ok {
		return cacheData
	}
	return
}

func (analy *analysisCache) Clean() {
	analy.lock.Lock()
	defer analy.lock.Unlock()
	analy.data = make(map[string][]*analysisModel)
}
