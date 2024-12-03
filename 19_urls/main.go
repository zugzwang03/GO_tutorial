package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://leetcode.com:3000/learn?course=reactjs"

func main() {
	fmt.Println("Welcome to handling urls in golang")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params is %T\n", qparams)

	fmt.Println(qparams["course"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "leetcode.com",
		Path: "/learn",
		RawPath: "user=shreya",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}