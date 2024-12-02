package main

import "fmt"

func main() {
	fmt.Println("Welcome to if and else in golang")

	logInCount := 10
	var result string

	if logInCount < 10 {
		result = "Regular user"
	} else if logInCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	var num int

	if num % 2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is NOT less than 10")
	}
}