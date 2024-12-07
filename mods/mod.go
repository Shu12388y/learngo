package main

import (
	"fmt"
	"github/mods/modules" 
)

func main() {
	var number int
	fmt.Println("Enter a number:")
	fmt.Scanf("%d", &number)
	// Call the function from the modules package
	modules.CheckNumber(number)
}
