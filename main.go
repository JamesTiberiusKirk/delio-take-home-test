package main

import (
	"delio-take-home-test/cli"
	finnhubwrapper "delio-take-home-test/finnhub_wrapper"
	"delio-take-home-test/stockchart"
)

func main() {
	flags := buildFlags()
	config := buildConfig()

	fc := initFinnhubApiClient(config.ApiKey)
	f := finnhubwrapper.NewFinnhubWrapper(fc)

	chart := stockchart.NewStockChart(f)

	c := cli.NewCli(chart)
	c.Run(flags.Stocks)
}
