package main

import "fmt"

func main()  {
	
	// define variables
	var name string = "Golang";
	fmt.Println(name);


	// infer
	var named = "Golang";
	fmt.Println(named);


	// shorthand syntax
	names := "Golang";
	fmt.Println(names);


	// define constant
	const age  =  12;
	fmt.Print(age)
	
	// const grouping

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Print(port,host)


}