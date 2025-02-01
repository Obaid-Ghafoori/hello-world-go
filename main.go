package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello World!")
	printHelloWorld()
	printUserData("John Doe", 30, "test#emia.com")
}

// write function to print hello world
func printHelloWorld() {
	fmt.Println("Hello World!")
}

// write fuction to print user data
func printUserData(name string, age int, email string) {
	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Email: ", email)
}
