package main

import (
	"github.com/HankWang95/test-demo/book"
	"fmt"
)

func main() {
	//book.ReadFile1_4()
	//demo.UitableDemo()
	//book.Image()
	//book.Fetch1_6()
	//book.WebHandler()
	//book.Echo4()
	//s := book.Basename("/ect/e/hi.go")

	s := book.BufComma("1234567890145678")
	fmt.Println(s)
}
