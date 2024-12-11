package main

import (
	"fmt"
)

func main() {

	var i int

	switch i {
	case 1:
		fmt.Println("Condition 1")
	case 2:
		fmt.Println("Condition 2")
	default:
		fmt.Println("Other")
	}



	whoIam := func (j interface{})  {
		switch j.(type){
		case int:
			fmt.Println("Integer")
		case string:
			fmt.Println("String")
		default:
			fmt.Println("Other")
		}
	}

	whoIam("String")

}