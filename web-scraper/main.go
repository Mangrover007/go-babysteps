package main

import (
	"net/http"
	"io"
	"strings"
	"golang.org/x/net/html"
	"fmt"
	"net/url"
	"errors"
)

const baseURL string = "https://scrape-me.dreamsofcode.io"

func getPage(pagelink *url.URL) []byte {
	page, err := http.Get(pagelink.String())
	if err != nil {
		panic(err)
	}

	if page.StatusCode >= http.StatusBadRequest {
		fmt.Printf("%s is a dead link\n", pagelink.String())
		return []byte("")
	}

	defer page.Body.Close()

	body, err := io.ReadAll(page.Body)
	if err != nil {
		panic(err)
	}

	return body
}

func parseAnchor(attributes []html.Attribute) (*url.URL, error) {
	for _, val := range attributes {
		if val.Key == "href" {
			if val.Val[0:1] == "/" {
				link, _ := url.Parse(baseURL + val.Val)
				return link, nil
			} else {
				link, _ := url.Parse(val.Val)
				return link, nil
			}
			break
		}
	}

	return nil, errors.New("no href in this wtf?")
}

func parsePage(pagelink *url.URL, linkchan chan<- *url.URL) {
	body := getPage(pagelink)

	r := strings.NewReader(string(body))
	z := html.NewTokenizer(r)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := z.Token()
			if t.Data == "a" {
				link, err := parseAnchor(t.Attr)
				if err != nil {
					return
				} else {
					linkchan <- link
				}
			}
			break
		}
	}
}

func main() {
	mainLink, _ := url.Parse(baseURL)
	mainHost := mainLink.Host

	visited := make(map[string]bool)
	
	linkchan := make(chan *url.URL)

	visited[mainLink.String()] = true
	go parsePage(mainLink, linkchan)

	for {
		link := <- linkchan
		if (link.Host == mainHost) && !visited[link.String()] {
			visited[link.String()] = true
			go parsePage(link, linkchan)
		}
	}
}

