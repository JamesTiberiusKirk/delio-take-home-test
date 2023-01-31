//go:generate mockgen -package stockchart -destination stock_chart_mock.go -source stock_chart.go

package stockchart

import (
	"fmt"

	"github.com/Finnhub-Stock-API/finnhub-go/v2"
)

type StockResult struct {
	StockName                string
	CurrentPrice             float32
	PreviousClosingPrice     float32
	Difference               float32
	TenXCurrentPrice         float32
	TenXPreviousClosingPrice float32
	TenXDifference           float32
}

type stockRepo interface {
	GetQuote(stock string) (*finnhub.Quote, error)
}

type StockChart struct {
	repo stockRepo
}

func NewStockChart(repo stockRepo) *StockChart {
	return &StockChart{
		repo: repo,
	}
}

func (c *StockChart) GetChart(stockName string) (*StockResult, error) {
	result := StockResult{}

	stockData, err := c.repo.GetQuote(stockName)
	if err != nil {
		return nil, fmt.Errorf("error fetching quote for stock %s: %v", stockName, err)
	}

	if stockData == nil {
		return nil, fmt.Errorf("no data for %s", stockName)
	}

	if stockData.C == nil || stockData.D == nil || stockData.Pc == nil {
		return nil, fmt.Errorf("missing data for %s", stockName)
	}

	txCurrentPriceCh := make(chan float32)
	txPrevousClosingPriceCh := make(chan float32)
	txDifferenceCh := make(chan float32)

	result.StockName = stockName

	// NOTE: Yes I know, concurrency here doesn't make sense...explanation in README
	result.CurrentPrice = stockData.GetC()
	go concurentTenX(txCurrentPriceCh, result.CurrentPrice)

	result.PreviousClosingPrice = stockData.GetPc()
	go concurentTenX(txPrevousClosingPriceCh, result.PreviousClosingPrice)

	result.Difference = stockData.GetD()
	go concurentTenX(txDifferenceCh, result.Difference)

	result.TenXCurrentPrice = <-txCurrentPriceCh
	result.TenXPreviousClosingPrice = <-txPrevousClosingPriceCh
	result.TenXDifference = <-txDifferenceCh

	return &result, nil
}

func concurentTenX(ch chan float32, val float32) {
	ch <- val * 10
}
