package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"

	// test := "Test String" // We have to use colon only during the initialization.
	// // var - variable
	// // card - name of the Variable
	// // string - only a String will ever be assigned to this variable

	// //Go is a statically typed programming language

	// fmt.Println(card)
	// fmt.Println(test)

	// test = "Test string changed"
	// fmt.Println(test)

	card := newCard()
	fmt.Println(card)

}

func newCard() string {
	return "Five of diamonds"
}
