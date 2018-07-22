package book

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"os"
)

// html 包内容
//type Node struct {
//	Type NodeType
//	Data string
//	Attr []Attribute
//	FirstChild, NextChild *Node
//
//}
//
//type NodeType int32
//
//type Attribute struct {
//	Key, Val string
//}
//
//const (
//	ErrorNode NodeType = iota
//	TextNode
//	DocumentNode
//	ElementNode
//	CommentNode
//	DoctypeNode
//)
//func Parse(r io.Reader) (*Node, error)  {
//
//}

var htmlCount = map[string]int{}

func ScanInnerURL(b []byte) {
	r := bytes.NewReader(b)
	//dec := base64.NewDecoder(base64.StdEncoding, buf)
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Fprintf(os.Stdout, "findlinks1: %v\n", link)
	}
}

// 函数会将n节点中的每一个链接添加到列表中
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func HTMLTree(b []byte) {
	r := bytes.NewReader(b)
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "HTMLTree: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
	fmt.Println("----------HTML Node Count----------")
	for k, v := range htmlCount {
		fmt.Printf("%s: \t %d\n", k, v)
	}
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		htmlCount[n.Data]++
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
	}
}
