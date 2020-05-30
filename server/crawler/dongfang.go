package crawler

import (
	"encoding/json"
	"sync"
	"taylors/model"
)

type dongFang struct {
	Rc   int `json:"rc"`
	Rt   int `json:"rt"`
	Svr  int `json:"svr"`
	Lt   int `json:"lt"`
	Full int `json:"full"`
	Data struct {
		Total int           `json:"total"`
		Diff  []model.Stock `json:"diff"`
	} `json:"data"`
}

var dongFangLock sync.Mutex

func (dongFangModel *dongFang) JsonToModel(jsonBody []byte) (dongFang *dongFang, err error) {
	dongFangLock.Lock()
	defer dongFangLock.Unlock()
	err = json.Unmarshal(jsonBody, &dongFang)
	return
}
