package main

import "fmt"

var jwtToken = 30000
const LogInToken string = "gibberish" // Public, so starts with caps

func main() {
	var username string = "shreya"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.4555678299
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and aliases
	var defaultValue int
	fmt.Println(defaultValue)
	fmt.Printf("Variable is of type: %T \n", defaultValue)

	var defaultValueInString string
	fmt.Println(defaultValueInString)
	fmt.Printf("Variable is of type: %T \n", defaultValueInString)

	// implicit type
	var website = "go.in"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	// no var style
	value := 300000.090
	fmt.Println(value)
	fmt.Printf("Variable is of type: %T \n", value)

	fmt.Println(LogInToken)
	fmt.Printf("The type is: %T", LogInToken)
}