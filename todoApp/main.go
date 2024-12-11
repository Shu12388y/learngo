package main

import (
	"fmt"
)


func main(){
		fmt.Print("Welcome to TODO app \n");
		for ;  ;  {
			var option int;
			fmt.Print("Enter options \n");
			fmt.Print("1. Add a todo \n");
			fmt.Print("2. Exit from the App \n")
			fmt.Scanf("%d",&option);
			if option == 1 {
				fmt.Print("\n Enter a new todo")
			}
			if option == 0 {
				break;
			}
		}
}