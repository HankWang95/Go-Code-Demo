package main

import (
	"github.com/HankWang95/test-demo/book"
)

func main() {
	//book.ReadFile1_3()
	//demo.UitableDemo()
	//book.Image()
	//book.Fetch1_6()
	//book.WebHandler()
	//book.Echo4()
	//s := book.Basename("/ect/e/hi.go")

	//s := book.BufComma("1234567890145678")
	//fmt.Println(s)
	//x := []int{1, 2, 3, 4}
	//y := book.AppendSlice(x, 3, 4, 4, 3, 3, 4)
	//fmt.Println(y)
	//book.Dedup()
	//book.WordFrequency()
	//values := []int{1,2,2,3,23,2,12,3,1,2}
	//book.TreeSort(values)
	//book.DoSearchIssue()
	//Spider()

	//book.DoTopology()
	//url := book.LoadUrl()
	//links, err := spider.Extract(url)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "err:%v", err)
	//
	//}
	//fmt.Println(links)
	//spider.CrawlerGo()
	//Spider()

	//url := book.LoadUrl()
	//spider.Outline(url)

	//number1 := book.Max(1,23,2,3,23,2,1111,2323,-123123)
	//number2 := book.Min()
	//fmt.Println(number1,number2)

	//url := book.LoadUrl()
	//node := book.GetAndParse(url)
	//book.SoleTitle(node)
	//fmt.Println(node.FirstChild.NextSibling.Data)
	//book.ElementsByTagName(node, "div", "h2", "p")

	//book.BigSlowOperation()

	book.PanicTest()
}

//
//func Spider() {
//	url := book.LoadUrl()
//	fmt.Printf("Loading %s ...\n", url)
//	resp, err := book.WaitForServer(url)
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "loading url fail! %s", err)
//	}
//
//	//body, err := book.BodyToByte(resp)
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "err :%v ", err)
//	}
//	//fmt.Println("---------Scan Inner URL-----------")
//	//book.ScanInnerURL(body)
//	fmt.Println("---------Scan the HTML tree struct---------")
//	spider.HTMLTree(resp)
//}
