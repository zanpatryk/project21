package deck

import (
	"strconv"
)

type Suit int

const (
	Clubs = iota
	Diamonds
	Hearts
	Spades
)

type Rank int

const (
	Jack  = 11
	Queen = 12
	King  = 13
	Ace   = 14
)

type Card struct {
	Suit Suit // 0-3, 0 - Clubs, 1 - Diamonds, 2 - Hearts, 3 - Spades
	Rank Rank // 2-14, 2-10, Jack, Queen, King, Ace
}

func (r Rank) String() string {
	switch r {
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Ace:
		return "A"
	default:
		return strconv.Itoa(int(r))
	}
}

func (s Suit) String() string {
	switch s {
	case Clubs:
		return "C"
	case Diamonds:
		return "D"
	case Hearts:
		return "H"
	case Spades:
		return "S"
	default:
		return "invalid suit"
	}
}

func (c Card) String() string {
	return c.Rank.String() + c.Suit.String()
}

func (c Card) Code() int {
	return int(c.Rank)*100 + int(c.Suit)
}
