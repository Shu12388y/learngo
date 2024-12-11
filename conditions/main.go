package main

import "fmt"

func main() {
	age := 1

	if age >= 18 {
		fmt.Println("Adult")
	} else if age >= 12 {
		fmt.Println("Teenage")
	} else{
		fmt.Println("Kid")
	}

}