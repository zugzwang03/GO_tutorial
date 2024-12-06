package main

import "fmt"

// model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) isEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to building apis in golang")
}
