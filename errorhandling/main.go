// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {

// 	res,err := http.Get("https://jsonplaceholder.typicode.com/posts");

// 	if err != nil{
// 		fmt.Println("Something went wrong")
// 		return
// 	}
// 	fmt.Println("Here is the result",res)

// }

package main

import "fmt"


func main(){

	// build a map
	hashedMap := map[string]int{
		"one":1,
		"two":2,
	}

	checkval,err := hashedMap["three"]


	if !err{
		fmt.Println("Value not found")
		return
	}

	fmt.Println(checkval)
}