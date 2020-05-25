package service

import (
	"context"
	"taylors/dao"
	"taylors/model"
	"taylors/module/crawler"
	"taylors_proto/taylors_stock"
)

type stockCommonService struct {
}

func (*stockCommonService) HistoryList(code string) (stockList []model.Stock, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), _OverTime)
	defer cancel()

	req := &taylors_stock.HistoryListReq{
		Code: code,
	}
	historyListRsp, err := crawler.Grpc_cli.HistoryList(ctx, req)
	if err != nil {
		return
	}

	stockList = Conv(historyListRsp.StockList)
	return
}

func (*stockCommonService) CodeDetail(codes []string) (stockList []model.Stock, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), _OverTime)
	defer cancel()

	req := &taylors_stock.CodeDetailReq{
		Code: codes,
	}

	historyListRsp, err := crawler.Grpc_cli.CodeDetail(ctx, req)
	if err != nil {
		return
	}

	stockList = Conv(historyListRsp.StockList)
	return
}

func (srv *stockCommonService) CheckOffday() (offday bool) {
	codes := []string{"SZ000001"}

	stockList, err := srv.CodeDetail(codes)
	if err != nil {
		return
	}
	if len(stockList) < 1 {
		return
	}

	stock := stockList[0]

	stockPOList, err := dao.StockDao.FindBySymbol(stock.Symbol)
	if err != nil {
		return
	}
	if len(stockPOList) == 0 {
		return
	}

	stockPO := stockPOList[0]

	if stock.Symbol != stockPO.Symbol {
		return
	}

	if stock.Percent != stockPO.Percent {
		return
	}

	if stock.Amount != stockPO.Amount {
		return
	}

	if stock.Volume != stockPO.Volume {
		return
	}

	offday = true

	return
}
