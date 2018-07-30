package goroutine

import (
	"fmt"
	"time"
)

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // 当res 准备好后关闭通道
}

type request struct {
	key      string
	response chan<- result // 客户端需要单个result
}

type Memo struct {
	requests chan request
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	fmt.Printf("Get %s time is %v\n", key, time.Now())
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// 对这个key的第一次请求
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // 调用f(key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// 等待数据准备完毕
	<-e.ready
	// 像客户端发送结果
	response <- e.res
}

func MemoFunc(key string) (interface{}, error) {
	var s time.Duration
	s = time.Duration(len(key))

	time.Sleep(time.Second * s)
	return key + time.Now().String(), nil
}

func DoMemo() {
	memo := New(MemoFunc)

	res1, err := memo.Get("hi")
	if err != nil {
		return
	}
	fmt.Println(res1)

	res2, err := memo.Get("hi")
	fmt.Println(res2)

	res3, err := memo.Get("hello")
	fmt.Println(res3)

	res4, err := memo.Get("hi")
	fmt.Println(res4)
}
