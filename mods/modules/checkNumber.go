package modules

import "fmt"


func CheckNumber(num int){
	if num%2 == 0{
		fmt.Println("The number is even")
	}else{
		fmt.Println(("The number is odd"))
	}

}