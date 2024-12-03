package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)

	fmt.Println("Result is: ", result)

	proResult, myMessage := proAdder(2, 5, 8, 7)

	fmt.Println("Result of proAdder is: ", proResult)
	fmt.Println("Pro message is: ", myMessage)
}

func adder(valOne int, valTwo int) int {
	return (valOne + valTwo)
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "pro adder function"
}

func greeter() {
	fmt.Println("Namastey from golang")
}

func greeterTwo() {
	fmt.Println("Another method")
}
