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

	txCurrentPriceCh := make(chan float32)
	txPrevousClosingPriceCh := make(chan float32)
	txDifferenceCh := make(chan float32)

	result.StockName = stockName

	if stockData.C != nil {
		result.CurrentPrice = *stockData.C
		go concurentTenX(txCurrentPriceCh, result.CurrentPrice)
	}

	if stockData.Pc != nil {
		result.PreviousClosingPrice = *stockData.Pc
		go concurentTenX(txPrevousClosingPriceCh, result.PreviousClosingPrice)
	}

	if stockData.D != nil {
		result.Difference = *stockData.D
		go concurentTenX(txDifferenceCh, result.Difference)
	}

	result.TenXCurrentPrice = <-txCurrentPriceCh
	result.TenXPreviousClosingPrice = <-txPrevousClosingPriceCh
	result.TenXDifference = <-txDifferenceCh

	return &result, nil
}

func concurentTenX(ch chan float32, val float32) {
	ch <- val * 10
}
