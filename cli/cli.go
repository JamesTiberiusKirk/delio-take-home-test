package cli

import (
	"delio-take-home-test/stockchart"
	"fmt"
	"sync"

	"github.com/fatih/color"
)

type ConcurrentStockData struct {
	Err  error
	Data *stockchart.StockResult
}

type Cli struct {
	chart *stockchart.StockChart
}

func NewCli(chart *stockchart.StockChart) *Cli {
	return &Cli{
		chart: chart,
	}
}

func (c *Cli) Run(stocks []string) {
	stocksData := map[string]*stockchart.StockResult{}

	wg := &sync.WaitGroup{}
	for _, s := range stocks {
		wg.Add(1)

		// Personally I do not like to embed channels into libraries and usually leave it
		//	up to the caller in case concurrency is not needed
		go func(stockName string) {
			defer wg.Done()
			stockData, err := c.chart.GetChart(stockName)
			stocksData[stockName] = nil
			if err != nil {
				// panic(err)
				return
			}

			stocksData[stockName] = stockData
		}(s)
	}

	wg.Wait()

	for k, v := range stocksData {
		if v == nil {
			fmt.Printf("%v\n", color.RedString("NO DATA FOR STOCK "+k))
			continue
		}

		fmt.Println("---------------------------------------------")
		fmt.Printf("Stock: \t\t\t%s x1\t\tx10\n", k)
		fmt.Printf("Current Price: \t\t%f\t%f\n", v.CurrentPrice, v.TenXCurrentPrice)
		fmt.Printf("Previous Closing price: %f\t%f\n", v.PreviousClosingPrice, v.TenXPreviousClosingPrice)

		if v.Difference < 0 {
			fmt.Printf("Loss: \t\t\t%v\n", color.RedString(fmt.Sprint(v.Difference, "\t\t", v.TenXDifference)))
		} else if v.Difference > 0 {
			fmt.Printf("Profit: \t\t%v\n", color.GreenString(fmt.Sprint(v.Difference, "\t\t", v.TenXDifference)))
		} else {
			fmt.Printf("Profit/Loss: \t\t%v\n", color.YellowString(fmt.Sprint(v.Difference, "\t\t", v.TenXDifference)))
		}
	}
	fmt.Println("---------------------------------------------")
}
