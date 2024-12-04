package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb in golang")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"
	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Response status code is: ", response.StatusCode)
	fmt.Println("Content Length of response is: ", response.ContentLength)

	content, _ := io.ReadAll(response.Body)
	// fmt.Println("Content is: ", string(content))

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println("Response in string is: ", responseString.String())
}
