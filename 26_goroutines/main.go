package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	go greeter("Hello")
	greeter("world")
	websitelist := []string{
		"https://google.com",
		"https://go.dev",
		"https://google.in",
		"https://fb.com",
		"https://github.com",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("something wrong with endpoint")
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
