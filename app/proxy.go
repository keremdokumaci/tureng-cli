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

type LanguageSet struct {
	SourceLanguageShortForm string
	DestLanguageShortForm   string
	SourceLanguage          string
	DestLanguage            string
}

func NewTurengProxy(language string) *TurengProxy {
	return &TurengProxy{
		language: language,
		client: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

func (p *TurengProxy) Query(word string, ls LanguageSet) ([]string, error) {
	url := fmt.Sprintf("https://tureng.com/en/%s/%s", p.language, word)

	resp, err := p.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocument(url)
	selection := doc.Find("table.table-hover.table-striped.searchResultsTable")

	var translations []string

	rows := selection.Children().First().Children()
	rows.EachWithBreak(func(i int, s *goquery.Selection) bool {
		if i > 2 {
			s.Children().Each(func(k int, child *goquery.Selection) {
				if child.HasClass(ls.SourceLanguageShortForm) || child.HasClass(ls.DestLanguageShortForm) {
					text := child.Children().First().Text()
					if strings.ToLower(text) != strings.ToLower(word) {
						translations = append(translations, text)
					}
				}
			})
		}

		return true
	})
	return translations, nil
}
