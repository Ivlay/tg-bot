package htmlParser

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type HtmlParser struct {
	url string
}

func NewParser(url string) *HtmlParser {
	return &HtmlParser{url: url}
}

func (p *HtmlParser) parse(idNode string) string {
	res, err := http.Get(p.url); if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatal("Close connection")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body); if err != nil {
		log.Fatal(err)
	}

	return doc.Find(idNode).Find("li").First().Text()
}

func (p *HtmlParser) GetPrice() string {
	return p.parse("#macbook_pro_14_2021_")
}
