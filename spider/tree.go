package spider

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
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

func ScanInnerURL(b []byte) *html.Node {
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
	return doc
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

//func HTMLTree(resp *http.Response) {
//	doc, err := html.Parse(resp.Body)
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "HTMLTree: %v\n", err)
//		os.Exit(1)
//	}
//
//}

func Outline(url string) error {
	var f1, f2 func(node *html.Node)
	var depth int
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)

	if err != nil {
		return err
	}
	f1 = func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}
	f2 = func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}

	//!+call
	forEachNode(doc, f1, f2)
	//!-call

	return nil
}

//func startElement(n *html.Node) {
//	if n.Type == html.ElementNode {
//		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
//		depth++
//	}
//}
//
//func endElement(n *html.Node) {
//	if n.Type == html.ElementNode {
//		depth--
//		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
//	}
//}
