package main

import "fmt"

func main() {
	fmt.Println("If else in go lang")
	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Reglar user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login counts"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	}else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less that 10")
	} else {
		fmt.Println("Num is NOT less that 10")
	}
}