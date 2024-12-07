package main

import "fmt"

// array in go
// func main()  {
// 	arr := [3]int{1,2,3}
// 	// fmt.Println(arr)
// 	for i := 0; i < 3; i++ {
// 		fmt.Println(arr[i])
// 	}
// }



func main() {
	// find the even numbers in the array and create a new array Earr
	arr  :=[5]int{1,2,3,4,5}
	earr := []int{}
	for i := 0; i < len(arr); i++ {
		if arr[i] % 2 == 0{
			earr = append(earr, arr[i])
		}
	}
	fmt.Println(earr)
}