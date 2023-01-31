package stockchart

import (
	"errors"
	"testing"

	"github.com/Finnhub-Stock-API/finnhub-go/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStockChart_SuccessfullyGetStockData(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("with full stock data", func(t *testing.T) {
		expected := &StockResult{
			StockName:                "test",
			CurrentPrice:             float32(10),
			PreviousClosingPrice:     float32(9),
			Difference:               float32(1),
			TenXCurrentPrice:         float32(100),
			TenXPreviousClosingPrice: float32(90),
			TenXDifference:           float32(10),
		}

		mockedQuote := &finnhub.Quote{}
		mockedQuote.SetC(10)
		mockedQuote.SetD(1)
		mockedQuote.SetPc(9)

		mockRepo := NewMockstockRepo(ctrl)
		mockRepo.EXPECT().
			GetQuote("test").
			Return(mockedQuote, nil)

		stockChart := NewStockChart(mockRepo)
		actual, err := stockChart.GetChart("test")
		require.NoError(t, err)

		assert.Equal(t, *expected, *actual)
	})
}

func TestStockChart_GetErrorOn(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("api error", func(t *testing.T) {
		expextedErr := errors.New("test error")

		mockRepo := NewMockstockRepo(ctrl)
		mockRepo.EXPECT().
			GetQuote("test").
			Return(nil, expextedErr)

		stockChart := NewStockChart(mockRepo)
		actual, err := stockChart.GetChart("test")
		require.Nil(t, actual)

		assert.ErrorContains(t, err, "test error")
	})
	t.Run("with empty stock data", func(t *testing.T) {
		mockRepo := NewMockstockRepo(ctrl)
		mockRepo.EXPECT().
			GetQuote("test").
			Return(nil, nil)

		stockChart := NewStockChart(mockRepo)
		actual, err := stockChart.GetChart("test")
		require.Nil(t, actual)

		assert.ErrorContains(t, err, "no data for test")
	})
	t.Run("with missing stock data", func(t *testing.T) {
		mockedQuote := &finnhub.Quote{}

		mockRepo := NewMockstockRepo(ctrl)
		mockRepo.EXPECT().
			GetQuote("test").
			Return(mockedQuote, nil)

		stockChart := NewStockChart(mockRepo)
		actual, err := stockChart.GetChart("test")
		require.Nil(t, actual)

		assert.ErrorContains(t, err, "missing data for test")
	})
}
