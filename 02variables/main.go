package main

import "fmt"

func main() {
	var username string = "Amal";
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n ",username)

	var isLoogedIn  bool = true;
	fmt.Println(isLoogedIn)
	fmt.Printf("variable is of type: %T \n ",isLoogedIn)

	var smallVal  uint8 = 255;
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n ",smallVal)

	var smallFloat  float64 = 255.0121333333333333;
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n ",smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n ",anotherVariable)
}