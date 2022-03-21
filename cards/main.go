package main

import "fmt"

func main()  {
	// var card string = "Ace of Spades"
	// ^^ long-winded version of the below
	// only use := on initialization of variable, not on reassignment
	card := newCard()
	// card = "Five of Diamonds"

	fmt.Println(card)
}
// functions can be declared in the same file as the main function, but if it will return value, need to specifically say what type it returns

func newCard() string {
	return "Five of Diamonds"
}

func secondNewCard() int {
	return 12
}