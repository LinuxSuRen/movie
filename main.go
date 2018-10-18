package main

import (
	"github.com/gocolly/colly"
	"strings"
)

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request){
		fmt.Println("visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error){
		log.Println("something wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("visited", r.URL)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement){
		link := e.Attr('href')
		
		if strings.HasPrefix('magnet:', link) {
			fmt.Println('magnet link', link)
		} else {
			e.Request.Visit(e.Request.AbsoluteURL(link))
		}
	})

	c.OnScraped(func(r *colly.Response){
		fmt.Println("finisedh", r.URL)
	})

	c.Visit("http://www.meijutt.com/")
}