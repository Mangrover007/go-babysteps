package main

import (
	"net/http"
	"io"
	"strings"
	"golang.org/x/net/html"
	"fmt"
	"net/url"
)

const baseURL string = "https://scrape-me.dreamsofcode.io"

func parsePage(pagelink *url.URL, linkchan chan<- *url.URL) {
	page, err := http.Get(pagelink.String())
	if err != nil {
		panic(err)
	}

	if page.StatusCode >= http.StatusBadRequest {
		fmt.Printf("%s is a dead link\n", pagelink.String())
		return
	}

	defer page.Body.Close()

	body, err := io.ReadAll(page.Body)
	if err != nil {
		panic(err)
	}

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
				for _, val := range t.Attr {
					if val.Key == "href" {
						if val.Val[0:1] == "/" {
							link, _ := url.Parse(baseURL + val.Val)
							linkchan <- link
						} else {
							link, _ := url.Parse(val.Val)
							linkchan <- link
						}
						break
					}
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
		// fmt.Println(link.String())
	}
}

