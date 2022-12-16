package htmlParser

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	macbook_pro_14_2021 = "#macbook_pro_14_2021_"
	macbook_pro_16_2021 = "#macbook_pro_16_2021_"
	mac_studio = "#mac_studio_"
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

	words := doc.Find(idNode).Find("li").Not(":last-of-type").Map(func(_ int, sel *goquery.Selection) string {
		return fmt.Sprintf("%s\n", sel.Text())
	})

	return strings.Join(words, "")
}

func (p *HtmlParser) GetPrice(idNode string) string {
	return p.parse(idNode)
}
