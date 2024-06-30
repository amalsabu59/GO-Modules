package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int

	// fmt.Println("value of pointer is ", ptr)

	myNumber := 23;

	var prt = &myNumber

	fmt.Println("value of pointer is ", prt);
	fmt.Println("value of pointer is ", *prt);


}