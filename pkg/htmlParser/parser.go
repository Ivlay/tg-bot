package htmlParser

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	tgbot "github.com/Ivlay/go-telegram-bot"
	"github.com/PuerkitoBio/goquery"
)

const (
	macbook_pro_14_2021 = "#macbook_pro_14_2021_"
	macbook_pro_16_2021 = "#macbook_pro_16_2021_"
	mac_studio          = "#mac_studio_"
)

type HtmlParser struct {
	url string
}

func NewParser(url string) *HtmlParser {
	return &HtmlParser{url: url}
}

func (p *HtmlParser) prepareProducts(sel *goquery.Selection, idNode string) []tgbot.Product {
	pp := make([]tgbot.Product, 0)
	sel.Find(idNode).Find("li").Not(":last-of-type").Each(func(_ int, sel *goquery.Selection) {
		arrStr := strings.Split(sel.Text(), " — ")
		productName := arrStr[0]
		pr := regexp.MustCompile("[0-9]+").FindString(strings.ReplaceAll(arrStr[1], " ", ""))
		if pr == "" {
			pr = "0"
		}
		price, err := strconv.ParseFloat(pr, 64)
		if err != nil {
			log.Fatal(err)
		}

		p := tgbot.Product{
			Price:   int(price),
			Title:   productName,
			PriceId: idNode,
		}

		pp = append(pp, p)
	})

	return pp
}

func (p *HtmlParser) PrepareProducts() []tgbot.Product {
	doc, err := p.getMainDoc()
	if err != nil {
		log.Fatal(err)
	}

	mac14 := p.prepareProducts(doc, macbook_pro_14_2021)
	mac16 := p.prepareProducts(doc, macbook_pro_16_2021)

	pp := append(mac14, mac16...)

	return pp
}

func (p *HtmlParser) getMainDoc() (*goquery.Selection, error) {
	res, err := http.Get(p.url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatal("Close connection")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	body := doc.Find("html body")

	return body, nil
}
