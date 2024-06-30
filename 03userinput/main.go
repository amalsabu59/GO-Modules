package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Enter the rate for our pizza:")

	// comma ok \\ error ok sytax
	input, _ := reader.ReadString('\n') 
	
	fmt.Println("Thanks fo rating,",input)
	fmt.Printf("Type of rating is %T",input)

}