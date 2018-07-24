package spider

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

var doc *html.Node

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err:%v", err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s:%s", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, f1, f2 func(n *html.Node)) {
	if f1 != nil {
		f1(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, f1, f2)
	}
	if f2 != nil {
		f2(n)
	}

}
