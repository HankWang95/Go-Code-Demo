package book

import (
	"net/http"
	"sync"
	"fmt"
	"log"
	"strconv"
)

var	mu sync.Mutex
var count  = 0

func WebHandler()  {
	http.HandleFunc("/count", counter)
	http.HandleFunc("/gif", handlerLissajous)
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:8999", nil))
}
// 引入 锁
func handler(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s %s %s\n",r.Method, r.URL, r.Proto)
	for k,v := range r.Header{
		fmt.Fprintf(w, "Header[%s] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err:= r.ParseForm(); err!=nil{
		log.Print(err)
	}
	for k,v := range r.Form{
		fmt.Fprintf(w, "Form[%q] = %q\n",k,v)
	}
}

func counter(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	fmt.Fprintf(w, "count:%d\n", count)
	mu.Unlock()

}

func handlerLissajous(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	cycles, _ := strconv.ParseFloat(r.Form.Get("cycles"), 10)
	lissajous(w,cycles)
}