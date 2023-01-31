package main

import finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"

func initFinnhubApiClient(apiKey string) *finnhub.DefaultApiService {
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", apiKey)
	return finnhub.NewAPIClient(cfg).DefaultApi
}
