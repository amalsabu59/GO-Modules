package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple","Banana"}
	fmt.Printf("Type of fruitlist %T \n",fruitList)

	fruitList = append(fruitList,"Mango","Banana",)
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScore := make([]int,4)
	
	highScore[0] = 234
	highScore[1] = 934
	highScore[2] = 534
	highScore[3] = 734
	// highScore[4] = 634

	highScore = append(highScore, 555,666,888)
	fmt.Println((highScore))

	sort.Ints(highScore)
	fmt.Println((highScore))

	// how to reave a vlaue form slice based on the index

	var courses = []string{"reactjs","javascript","swift","pthon","ruby"}
	fmt.Println(courses)
	var index int = 2;
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)
}