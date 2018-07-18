package book

import (
	"strings"
	"bytes"
	"fmt"
)

func Basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

// 函数接收 整数字符串 "12345"， 返回"12,345"
func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return Comma(s[:n-3]) + "," + s[n-3:]
}

func BufComma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	fmt.Println(s[1])
	for i:=0 ; i<n; i++{
		fmt.Println(i)
		buf.WriteByte(s[i])
		if n!=i+1 && (n-i-1)%3 == 0{
			buf.WriteByte(',')
		}
	}

	return buf.String()
}
