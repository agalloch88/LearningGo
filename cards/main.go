package main

import "fmt"

func main()  {
	// var card string = "Ace of Spades"
	// ^^ long-winded version of the below
	// only use := on initialization of variable, not on reassignment
	card := "Ace of Spades"
	card = "Five of Diamonds"
	
	fmt.Println(card)
}