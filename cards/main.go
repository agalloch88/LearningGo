package main

// import "fmt"

func main()  {
	// var card string = "Ace of Spades"
	// ^^ long-winded version of the below
	// only use := on initialization of variable, not on reassignment
	// card := newCard()
	// // card = "Five of Diamonds"

	// fmt.Println(card)

	// how to declare a slice (aka dynamic array) and add new values to it
	// append does not modify existing slice, but rather returns a new slice and reassigns the variable

	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// // how to set up for loop in go

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// greeting := "Hi There!"
	// fmt.Println([]byte(greeting))
	cards := newDeck()
	cards.saveToFile("my_cards")
}
// functions can be declared in the same file as the main function, but if it will return value, need to specifically say what type it returns

// func newCard() string {
// 	return "Five of Diamonds"
// }

// func secondNewCard() int {
// 	return 12
// }

// slice range syntax:
// fruits[startIndexIncluding : upToNotIncluding]
// fruits[0:2], for example, returns indexes 0 and 1, but not 2