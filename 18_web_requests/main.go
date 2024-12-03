package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://leetcode.com/u/Zugzwang03/"

func main() {
	fmt.Println("Welcome to web requests")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of the response is: %T \n", response)
	defer response.Body.Close() // caller's responsibility to close the connection

	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(dataBytes)
	fmt.Println("The content of response is: ", content)
}