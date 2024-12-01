package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Orange"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Length of fruit list is: ", len(fruitList))

	vegList := [5]string{"potato", "beans", "mushroom"}
	fmt.Println("Veg list is: ", vegList)
	fmt.Println("Length of veg list is: ", len(vegList))
}