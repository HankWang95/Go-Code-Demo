package book

import (
	"fmt"
	"golang.org/x/net/html"
	"math"
	"net/http"
)

func Max(vals ...int) int {
	var max int
	max = math.MinInt32
	if len(vals) == 0 {
		fmt.Println("please input some number!")
		return 0
	}
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func Min(vals ...int) int {
	var min int
	min = math.MaxInt32
	if len(vals) == 0 {
		fmt.Println("please input some number!")
		return 0
	}
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func GetAndParse(url string) *html.Node {
	doc, err := http.Get(url)
	if err != nil {
		fmt.Println("err")
	}
	node, err := html.Parse(doc.Body)
	if err != nil {
		fmt.Println(err)
	}
	return node
}

func ElementsByTagName(n *html.Node, name ...string) []*html.Node {
	var NodeList []*html.Node
	var strList = make(map[string]bool)
	for _, val := range name {
		if !strList[val] {
			strList[val] = true
		}
	}
	fmt.Println(strList)
	f := func(n *html.Node) {
		if n.Type == html.ElementNode {
			if strList[n.Data] {
				NodeList = append(NodeList, n)
				fmt.Println(n.Data)
			}
		}
	}

	visitNode(n, f)

	return NodeList
}

func visitNode(n *html.Node, f func(n *html.Node)) {
	f(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visitNode(c, f)
	}
}
