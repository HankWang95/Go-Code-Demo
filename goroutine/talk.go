package goroutine

import (
	"fmt"
	"time"
)

var count int
var msg chan string

type people struct {
	Name string
}

func (a *people) SendMsg(b *people, m chan string) {
	//fmt.Println("send msg before 2s ")
	//time.Sleep(2*time.Second)
	m <- a.Name // 发送也会阻塞，等待接收完才会释放
	fmt.Println("send ok")
}

func (a *people) ReciveMsg(m chan string) {
	//fmt.Println("waitting msg...")
	c := <-m
	fmt.Println("name:", a.Name, "recive ok msg:", c)
	count++

}

func NewMsg() chan string {
	msg := make(chan string)
	return msg
}

func NewPeople() (a, b *people) {
	a = &people{Name: "A"}
	b = &people{Name: "B"}
	return a, b
}

func DoMsg() {
	a, b := NewPeople()
	//c, d := NewPeople()// 一条信道的值不能够重复提取：多对象接收时 先到先得
	msg = NewMsg()
	for i := 0; i < 10; i++ {
		go a.SendMsg(b, msg)
		go b.ReciveMsg(msg)
		//go c.ReciveMsg(msg)
		//go d.ReciveMsg(msg)
	}

	time.Sleep(1000000000)
	fmt.Println(count)

	//b.SendMsg(a,msg)
}
