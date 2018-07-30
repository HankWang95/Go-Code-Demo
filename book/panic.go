package book

import (
	"fmt"
	"golang.org/x/net/html"
)

func SoleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
			fmt.Println("未发生宕机")
			//未发生宕机
		case bailout{}:
			fmt.Println("预期的宕机")
			//预期的宕机
			err = fmt.Errorf("multiple title elements")
		default:
			//未预期的宕机：继续宕机
			panic(p)
		}
	}()

	visitNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			if title != "" {
				fmt.Println(title)
				panic(bailout{}) // 多个标题元素引发宕机
			}
			title = n.FirstChild.Data
		}
	})
	if title == "" {
		fmt.Println("err")
		return "", fmt.Errorf("no title element")
	}
	fmt.Println(title)
	return title, nil

}

// 5.19 使用Panic 和 recover 写一个没有return 却能返回一个非零的值的函数

func PanicTest() {
	defer func() {
		recover()
	}()
	panic("非零值哦")
}
