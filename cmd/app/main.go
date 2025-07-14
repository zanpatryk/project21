package main

import (
	"fmt"

	"github.com/zanpatryk/project21/internal/deck"
)

func main() {
	deck := deck.NewDeck()
	deck.Shuffle()
	fmt.Println(deck)
}
