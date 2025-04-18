package main

import (
	"fmt"
	"time"
)

func main() {
	var count int64 = 0
	startTime := time.Now()                      // Initialize startTime
	for time.Since(startTime) < 10*time.Second { // Loop until 1 minute has passed
		// fmt.Println("GoLang: looped", count)
		count++
		// time.Sleep(1 * time.Second) // Sleep for 1 second to prevent overloading the CPU
	}
	fmt.Println("Done ", count)
}
