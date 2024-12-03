package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "Content for the file"
	file, err := os.Create("./myFile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./myFile.txt")
}

func readFile(fileName string) {
	dataByte, err := os.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println("Text data inside the file is \n", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}