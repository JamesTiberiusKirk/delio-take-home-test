package finnhubwrapper

import (
	"context"
	"fmt"

	"github.com/Finnhub-Stock-API/finnhub-go/v2"
)

// I'm unsure how to mock finnhub.ApiQuoteRequest
type FinnhubInterface interface {
	Quote(ctx context.Context) finnhub.ApiQuoteRequest
}

func NewFinnhubWrapper(apiClient FinnhubInterface) *FinnhubWapper {
	return &FinnhubWapper{
		fc: apiClient,
	}
}

type FinnhubWapper struct {
	fc FinnhubInterface
}

func (c *FinnhubWapper) GetQuote(s string) (*finnhub.Quote, error) {
	// TODO: TODOvsBackground?
	res, _, err := c.fc.Quote(context.Background()).
		Symbol(s).
		Execute()

	if err != nil {
		return nil, fmt.Errorf("unable to get current stock price for stock %s: %e", s, err)
	}

	return &res, nil
}
