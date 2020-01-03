package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	STOCK_CODE := "BAC"
	c := colly.NewCollector()
	// Find stock price
	c.OnHTML(".wsod_last", func(e *colly.HTMLElement) {
		fmt.Println(e.ChildText("span"))
	})
	c.OnHTML(".wsod_change", func(e *colly.HTMLElement) {
		fmt.Println(e.ChildText("span.posData"))
	})

	c.Visit("https://money.cnn.com/quote/forecast/forecast.html?symb=" + STOCK_CODE)
}
