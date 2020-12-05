// server1 是一个迷你的回声和计数器服务器
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mux sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counterHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

// handler 处理程序回显请求的URL 的路径部分
func handler(w http.ResponseWriter, r *http.Request) {
	mux.Lock()
	defer mux.Unlock()

	count++
	fmt.Fprintf(w, "URL.Path= %q\n", r.URL.Path)
}

// counterHandler 回显目前为止调用的次数
func counterHandler(w http.ResponseWriter, r *http.Request) {
	mux.Lock()
	defer mux.Unlock()

	fmt.Fprintf(w, "Count %d\n", count)
}
