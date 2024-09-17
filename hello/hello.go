package main

import(
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a greetings message and print it
	message := greetings.Hello("Shreya")
	fmt.Println(message)
}