package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	greeterTwo()

	result := adder(3,5);
	fmt.Println("Result is:",result)

	proResult := proAdder(1,2,3,4,6)

	fmt.Println("Pro result is:",proResult)

}

func adder(valOne int,valTwo int) int{
	return valOne + valTwo
}
func proAdder(values ...int) int {
	total := 0

	for _,value := range values {
		total += value
	}
return total
}
func greeterTwo()  {
		fmt.Println("Another method")
	}
func greeter() {
	fmt.Println("Namstey from golang")
}

