package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb in golang")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
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

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	// fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename":"let's go with golang",
			"price": 0,
			"platform":"anywebsite"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	// formdata

	data := url.Values{}
	data.Add("firstname", "shreya")
	data.Add("lastname", "sarkar")
	data.Add("email", "shreya@gmail.com")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}