package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	log.Printf("Application started at: %v", time.Now())
}

// func badpractice() {

// 	fmt.Println("Hello World from the bad practice!")

// 	// Adding logs
// 	log.Printf("Application started at: %v", time.Now())
// }
