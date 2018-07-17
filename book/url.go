package book

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"io"
)

func Fetch1_5() {
	for _, url := range os.Args[1:] {
		if url[:7] != "http://" {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v \n", err)
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v \n", err)
		}
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s:%v \n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)

	}
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
		ch<- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s",secs,nbytes,url)


}



