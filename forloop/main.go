package main

import "fmt"

func main() {

	// implement while loop in golang as there is not while loop in golang
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// infinte loop
	for {
		fmt.Println((1))
	}

	// classic for loop
	// for j := 0; j < 4; j++ {
	// 	fmt.Println(j)
	// }

	// using range
	// for i := range 11{
	// 	fmt.Println(i)
	// }

}
