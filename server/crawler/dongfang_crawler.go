package crawler

import (
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"taylors/global"
	"taylors/model"
	"time"
)

type dongFangCrawler struct {
	dongFangModel *dongFang
	url           string
	topUrl        string
	sleepTime     int
	env           string // dev stage prod
}

var (
	once               sync.Once
	dongFangCrawlerObj *dongFangCrawler
)

func NewDongFangCrawler() *dongFangCrawler {
	once.Do(func() {
		if !strings.Contains(global.GVA_CONFIG.Crawler.Url, "5000") {
			panic("Crawler.Url 配置错误")
		}
		dongFangCrawlerObj = &dongFangCrawler{
			dongFangModel: &dongFang{},
			url:           global.GVA_CONFIG.Crawler.Url,
			topUrl:        strings.ReplaceAll(global.GVA_CONFIG.Crawler.Url, "5000", "100"),
			sleepTime:     global.GVA_CONFIG.Crawler.Sleep,
			env:           global.GVA_CONFIG.Crawler.Env,
		}
	})
	return dongFangCrawlerObj
}

func (crawler *dongFangCrawler) get(url string) (body []byte, err error) {
	rsp, err := http.Get(url)
	if err != nil {
		return
	}
	rspBody := rsp.Body
	defer rspBody.Close()

	bys, err := ioutil.ReadAll(rspBody)
	if err != nil {
		return
	}

	body = bys
	return
}

func (crawler *dongFangCrawler) checkTime() (flag bool) {
	if crawler.env == "dev" {
		return true
	}

	now := time.Now()
	weekday := now.Weekday()
	if weekday == time.Sunday || weekday == time.Saturday {
		return false
	}

	forenoonStart := time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, now.Location())
	forenoonEnd := time.Date(now.Year(), now.Month(), now.Day(), 11, 40, 0, 0, now.Location())
	if now.After(forenoonStart) && now.Before(forenoonEnd) {
		return true
	}

	afternoonStart := time.Date(now.Year(), now.Month(), now.Day(), 12, 50, 0, 0, now.Location())
	afternoonEnd := time.Date(now.Year(), now.Month(), now.Day(), 15, 10, 0, 0, now.Location())
	if now.After(afternoonStart) && now.Before(afternoonEnd) {
		return true
	}

	return false
}

func (crawler *dongFangCrawler) Loop() {
	instence().store(crawler.process())
	for {
		if crawler.checkTime() {
			instence().store(crawler.process())
		}
		time.Sleep(time.Second * time.Duration(crawler.sleepTime))
	}
}

func (crawler *dongFangCrawler) process() (dongFang *dongFang) {
	body, err := crawler.get(crawler.url)
	if err != nil {
		return
	}
	dongFang, err = crawler.dongFangModel.JsonToModel(body)
	return
}

func (crawler *dongFangCrawler) processTop() (dongFang *dongFang) {
	body, err := crawler.get(crawler.topUrl)
	if err != nil {
		return
	}
	dongFang, err = crawler.dongFangModel.JsonToModel(body)
	return
}

func (crawler *dongFangCrawler) Top() (marketList []model.Stock) {
	dongFangTop := crawler.processTop()
	dongFangCodes := make([]string, 0)
	for _, dongFang := range dongFangTop.Data.Diff {
		dongFangCodes = append(dongFangCodes, dongFang.Code)
	}
	marketList = instence().obtain(dongFangCodes)
	return
}

func (crawler *dongFangCrawler) Monitor(codes []string) (marketList []model.Stock) {
	marketList = instence().obtain(codes)
	return
}

func (crawler *dongFangCrawler) All() (marketList []model.Stock) {
	marketList = instence().obtainAll()
	return
}
