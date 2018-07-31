package goroutine

import (
	"io"
	"log"
	"net"
	"time"
)

func DoClock() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go _handleConn(conn)

	}
}

func _handleConn(c net.Conn) {
	defer c.Close()
	for i := 1; i < 5; i++ {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
