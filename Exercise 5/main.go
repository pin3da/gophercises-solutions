package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"./parser"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	tmp := flag.String("start", "https://pin3da.github.io", "url that you want to build a sitemap for")
	maxDepth := flag.Int("depth", 10, "max depth of the pages")
	flag.Parse()

	startPage := *tmp

	if strings.HasSuffix(startPage, "/") {
		startPage = startPage[:len(startPage)-1]
	}

	pages := bfs(startPage, startPage, *maxDepth)

	toXML := urlset{
		Xmlns: xmlns,
	}
	for _, page := range pages {
		toXML.Urls = append(toXML.Urls, loc{page})
	}

	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(toXML); err != nil {
		panic(err)
	}
	fmt.Println()
}

func bfs(page, baseURL string, maxDepth int) []string {
	seen := make(map[string]struct{})
	seen[page] = struct{}{}
	var q, next []string
	q = append(q, page)
	for d := 0; d < maxDepth && len(q) > 0; d++ {
		for _, p := range q {
			all := getAll(p, baseURL)
			for _, link := range all {
				if _, ok := seen[link]; !ok {
					seen[link] = struct{}{}
					if strings.HasPrefix(link, baseURL) {
						next = append(next, link)
					}
				}
			}
		}
		q = next
		next = []string{}
	}

	ans := make([]string, 0, len(seen))

	for link := range seen {
		if strings.HasPrefix(link, baseURL) {
			ans = append(ans, link)
		}
	}

	return ans
}

func getAll(page, baseURL string) []string {
	resp, err := http.Get(page)
	if err != nil {
		return []string{}
	}
	defer resp.Body.Close()
	links, err := parser.ParseLinks(resp.Body)
	var ans []string
	for _, link := range links {
		ans = append(ans, normalizeURL(link.Href, baseURL))
	}
	return ans
}

func normalizeURL(href, baseURL string) string {
	if strings.HasPrefix(href, "./") {
		href = href[1:]
	}
	if strings.HasPrefix(href, "//") {
		href = "https:" + href
	}
	if strings.HasPrefix(href, "/") {
		href = baseURL + href
	}
	return href
}
