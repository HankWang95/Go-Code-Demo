package book

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)

func Fetch1_5()  {
	for _, url := range os.Args[1:]{
		if url[:7] != "http://"{
			url = "http://"+url
		}
		resp, err := http.Get(url)
		if err != nil{
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
