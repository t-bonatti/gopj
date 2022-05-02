package rfb

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const rfbBaseUrl = "http://200.152.38.155"

type RfbFiles struct {
	UpdatedAt time.Time
	FilesUrl  []string
}

type PageParser struct {
	client  *http.Client
	baseURL string
}

func NewPageParser() PageParser {
	return PageParser{client: http.DefaultClient, baseURL: rfbBaseUrl}
}

func (p *PageParser) GetRfbFiles() (rfbFiles RfbFiles, err error) {

	res, err := p.client.Get(p.baseURL + "/CNPJ")
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprint("status code error: ", res.Status))
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	doc.Find("td").Each(func(i int, s *goquery.Selection) {
		fileName, exists := s.Find("a").Attr("href")
		if exists && strings.Contains(fileName, ".zip") {
			rfbFiles.FilesUrl = append(rfbFiles.FilesUrl, p.baseURL+"/CNPJ/"+fileName)
		}
	})

	d := (24 * time.Hour)
	doc.Find("body > table > tbody > tr:nth-child(4) > td:nth-child(3)").Each(func(i int, s *goquery.Selection) {
		time, err := time.Parse("2006-01-02 15:04", strings.TrimSpace(s.Text()))
		fmt.Println(s.Text())
		if err != nil {
			fmt.Println(err)
			return
		}

		rfbFiles.UpdatedAt = time.Truncate(d)
	})

	return
}
