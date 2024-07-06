package main

import "fmt"

func main() {
	fmt.Println("Strcts in go lang")
	//no inherictance in golang; No super or parent 

	hitesh := User {
		"hitesh","hitesh@go.dev",true,16,
	}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v \n ",hitesh)
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}

