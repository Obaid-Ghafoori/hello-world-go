package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	log.Printf("Application started at: %v", LoginTime("admin"))
}

func LoginTime(username string) string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Printf("User %s logged in at %s\n", username, formattedTime)
	return formattedTime
}
