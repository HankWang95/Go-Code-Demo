package goroutine

import (
	"fmt"
	"github.com/HankWang95/pool4go"
	"io"
	"log"
	"net"
	"os"
)

func DoNetcat() {
	//ltcpaddr, err := net.ResolveTCPAddr("tcp", "localhost:8000")
	atcpaddr, err := net.ResolveTCPAddr("tcp", "localhost:8000")
	pool, err := pool4go.NewGPool(10, 30, func() (net.Conn, error) {
		return net.DialTCP("tcp", nil, atcpaddr)
	})

	//conn, err := net.DialTCP("tcp", nil, atcpaddr)
	conn, err := pool.Get()
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Print("down")
		done <- struct{}{} // 指示主 goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done
	fmt.Println("ok")
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
