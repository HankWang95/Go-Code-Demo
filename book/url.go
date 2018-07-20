package book

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func LoadUrl() string {
	url := os.Args[1]
	if url[:7] != "http://" {
		url = "http://" + url
	}
	//resp, err := http.Get(url)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "fetch: %v \n", err)
	//}
	//b, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "fetch: %v \n", err)
	//}
	//defer resp.Body.Close()
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "fetch: reading %s:%v \n", url, err)
	//	os.Exit(1)
	//}
	//fmt.Printf("%s", b)
	return url
}

// 引入chan概念，并发执行
func Fetch1_6() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		if url[:7] != "http://" {
			url = "http://" + url
		}
		go fetchGo(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchGo(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}

func WaitForServer(url string) (*http.Response, error) {
	const timeout = 1 * time.Minute
	deadLine := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadLine); tries++ {
		resp, err := http.Get(url)
		if err == nil {
			return resp, nil
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // 指数退避策略
	}
	return nil, fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func BodyToByte(resp *http.Response) ([]byte, error) {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read resp.body fail %s ", err)
	}
	return b, nil
}
