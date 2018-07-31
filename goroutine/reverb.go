package goroutine

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func HandelConn(c *net.TCPConn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	var abort = make(chan string)

	go func() {
		defer wg.Add(1)
		for {
			select {
			case <-time.Tick(10 * time.Second):
				c.CloseWrite()
			case text := <-abort:
				go echo(c, text, 1*time.Second)
			}
		}
	}()
	for input.Scan() {
		abort <- input.Text()
	}

}

func EchoConn() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:8000")
	listener, _ := net.ListenTCP("tcp", tcpAddr)
	for {
		conn, _ := listener.AcceptTCP()

		go HandelConn(conn)

	}
}
