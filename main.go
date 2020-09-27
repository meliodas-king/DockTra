package main

import "fmt"

func main() {
	wow := Eheh()
	fmt.Println(wow)
	cards := newDeck()
	cards.shuffle()
	cards.print()

}
