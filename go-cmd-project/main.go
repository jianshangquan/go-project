package main

import (
	"fmt"
	"os"
)





func main() {
	fmt.Println("Hello world");
	a, b := os.Getwd()
	if b == nil {
		fmt.Println(a)
	}	
}