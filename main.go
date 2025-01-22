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

func badpractice() {
	// Unused variable
	// deadVariable := "This variable is never used"

	// Active code
	fmt.Println("Hello World!")

	// Adding logs
	log.Printf("Application started at: %v", time.Now())

	// Commented out code (should be detected)
	/*
		if true {
			log.Println("This code is commented out")
		}
	*/
}
