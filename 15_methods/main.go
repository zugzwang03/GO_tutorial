package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in golang")

	// no inheritance in golang
	// no super or parent in golang
	// no classes in golang
	// not object oriented

	shreya := User{"Shreya", "shreya@gmail.com", true, 16}
	fmt.Println(shreya)
	fmt.Printf("shreya details are: %+v\n", shreya)
	fmt.Printf("Name is %v and email is %v.\n", shreya.Name, shreya.Email)
	shreya.GetStatus()
	shreya.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", shreya.Name, shreya.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "shreyasarkar@gmail.com"
	fmt.Println("Email of this user is: ", u.Email)
}