package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
