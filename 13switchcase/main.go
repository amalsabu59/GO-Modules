package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is",diceNumber)

	switch diceNumber {
	case 1 :
		fmt.Println("Dice value is 1 and you can open")
	case 2 :
		fmt.Println("You can move 2 Spot")	
	case 3 :
		fmt.Println("You can move 3 Spot")	
		fallthrough
	case 4 :
		fmt.Println("You can move 4 Spot")	
	case 5 :
		fmt.Println("You can move 5 Spot")	
	case 6 :
		fmt.Println("You can move 6 Spot and roll the dice again")
	default:
		fmt.Println("unknown")	
	}

}