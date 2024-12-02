package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Saturday"}

	fmt.Println(days)

	for d:=0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i:= range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v \n", index, day)
	}

	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 5 {
			break
		}
		if rougueValue == 4 {
			rougueValue++
			continue
		}
		if rougueValue == 2 {
			goto twoGotten
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

twoGotten:
	fmt.Println("Got two!")

}
