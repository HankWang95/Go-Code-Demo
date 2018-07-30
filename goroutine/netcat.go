package goroutine

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func DoNetcat() {
	//ltcpaddr, err := net.ResolveTCPAddr("tcp", "localhost:8000")
	atcpaddr, err := net.ResolveTCPAddr("tcp", "localhost:8000")
	conn, err := net.DialTCP("tcp", nil, atcpaddr)
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
	conn.CloseWrite()
	<-done
	fmt.Println("ok")
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
