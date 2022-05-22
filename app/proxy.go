package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type TurengProxy struct {
	client   *http.Client
	language string
}

func NewTurengProxy(language string) *TurengProxy {
	return &TurengProxy{
		language: language,
		client: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

func (p *TurengProxy) Query(word string) (string, error) {
	url := fmt.Sprintf("https://tureng.com/en/%s/%s", p.language, word)

	resp, err := p.client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	var response []string

	doc, err := goquery.NewDocument(url)
	selection := doc.Find("table.table-hover.table-striped.searchResultsTable")

	rows := selection.Children().First().Children()
	rows.EachWithBreak(func(i int, s *goquery.Selection) bool {
		if i > 3 {
			s.Children().Each(func(k int, child *goquery.Selection) {
				if k == 3 {
					response = append(response, child.Text())
				}
			})
		}

		if len(response) == 5 {
			return false
		}

		return true
	})
	return strings.Join(response, ","), nil
}
