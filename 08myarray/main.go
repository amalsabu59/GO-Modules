package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in go langs")

	var fruitList [4]string 
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[2] = "Peach"


	fmt.Println("fruit list is:" , fruitList)
	fmt.Println("fruit list is:" , len(fruitList))

	var vegList = [3]string {"potato","beans","mush"}
	fmt.Println("Vegyl list :", vegList)
	fmt.Println("Vegyl list :", len(vegList))

}