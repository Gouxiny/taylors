package initialize

import "taylors/crawler"

func Crawler() {
	go crawler.NewDongFangCrawler().Loop()
}
