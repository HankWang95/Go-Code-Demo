package pool

import (
	"fmt"
	"github.com/HankWang95/pool4go"
	"log"
	"net"
	"sync"
	"os"
	"time"
)

func init() {
	// used for factory function
	//go simpleTCPServer()
	//time.Sleep(time.Millisecond * 300) // wait until tcp server has been settled

}

func simpleTCPServer() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}


		//go func() {
		//	buffer := make([]byte, 256)
		//	conn.Read(buffer)
		//}()
		fmt.Fprint(os.Stdout, conn)
	}
}

func Demo4pool() {

	p, _ := pool4go.NewGPool(10, 30, func() (net.Conn, error) {
		return net.Dial("tcp", "localhost:8000")
	})
	defer p.Close()

	var wg sync.WaitGroup
	conns := make([]net.Conn, 31)
	for i := 0; i < 31; i++ {
		if i == 30{
			time.Sleep(2*time.Second)
			conn := conns[29]
			conn.Close()
		}

		if i == 8 {
			fmt.Println("--------测试连接池等待剩余----------")
			wg.Wait()
			fmt.Printf("In pool conner amount: %d\n", p.Len())
		}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			conn, err := p.Get()
			if err != nil{
				fmt.Println(err)
			}
			//msg := "hello"
			//fmt.Fprint(conn, msg)
			conns[i] = conn
			if err != nil {
				fmt.Errorf("Get error. ")
			}
		}(i)
	}
	wg.Wait()

	// now put them all back
	for i, conn := range conns {
		if i > 20 && i < 25 {
			func(i int) {
				_, err := p.Get()
				if err != nil {
					fmt.Errorf("Get error. ")
				}
				fmt.Printf("get again, amout : %d\n", p.Len())
			}(i)
			continue

		}
		conn.Close()
	}

	if p.Len() != 30 {
		fmt.Errorf("Put error len. Expecting %d, got %d ",
			1, p.Len())
	}

	conn, _ := p.Get()
	p.Close() // close pool

	conn.Close() // try to put into a full pool
	if p.Len() != 0 {
		fmt.Errorf("Put error. Closed pool shouldn't allow to put connections. ")
	}

}
