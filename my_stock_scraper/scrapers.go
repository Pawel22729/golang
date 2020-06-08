package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func get_stock(stockCode string) string {
	stockBaseUrl := "https://money.cnn.com/quote/forecast/forecast.html?symb="
	c := colly.NewCollector()
	// Find stock price
	c.OnHTML(".wsod_last", func(e *colly.HTMLElement) {
		res := e.ChildText("span")
		return res
	})
	c.OnHTML(".wsod_change", func(e *colly.HTMLElement) {
		res1 := e.ChildText("span.posData")
		return fmt.Sprint(res1)
	})

	c.Visit(stockBaseUrl + stockCode)
	resultString := fmt.Sprint(res1)
		
	return resultString
}
