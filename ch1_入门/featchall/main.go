// featchall 并发获取 URL 并报告它们的时间和大小
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go featch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // 从通道 ch 接收
	}

	fmt.Printf("%2.fs elapsed\n", time.Since(start).Seconds())
}

func featch(url string, ch chan<- string) {
	start := time.Now()
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	sec := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2.fs %7d %s", sec, nbytes, url)
}
