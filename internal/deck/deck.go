package deck

import (
	"math/rand"
)

type Deck []Card

func NewDeck() Deck {
	ranks := []Rank{2, 3, 4, 5, 6, 7, 8, 9, 10, Jack, Queen, King, Ace}
	suits := []Suit{Clubs, Diamonds, Hearts, Spades}

	deck := make(Deck, 0, 52)

	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, Card{Rank: rank, Suit: suit})
		}
	}

	return deck
}

func (d Deck) Shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
