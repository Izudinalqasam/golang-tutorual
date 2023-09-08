package goroutine

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s -> error: %s", url, err)
		return
	}

	defer resp.Body.Close()
	cType := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, cType)
}

func returnTypeCon(url string, out chan string) {
	resp, err := http.Get(url)
	if err != nil {
		out <- fmt.Sprintf("%s -> error: %s", url, err)
		return
	}

	defer resp.Body.Close()
	cType := resp.Header.Get("content-type")
	out <- fmt.Sprintf("%s -> %s\n", url, cType)
}

func siteSerial(urls []string) {
	for _, url := range urls {
		returnType(url)
	}
}

func sitesConcurrent(urls []string) {
	var wg sync.WaitGroup
	ch := make(chan string)
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			returnTypeCon(url, ch)
			wg.Done()
		}(url)
	}

	for range urls { // Run number of URLs times
		out := <-ch
		fmt.Println(out)
	}

	wg.Wait()
}

func MainGoroutineLearning() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/ip",
	}

	start := time.Now()
	siteSerial(urls)
	fmt.Println(time.Since(start))

	start = time.Now()
	sitesConcurrent(urls)
	fmt.Println(time.Since(start))
}
