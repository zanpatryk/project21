package deck

const (
	Clubs = iota
	Diamonds
	Hearts
	Spades
)

const (
	Jack  = 11
	Queen = 12
	King  = 13
	Ace   = 14
)

type Card struct {
	Suit int // 0-3, 0 - Clubs, 1 - Diamonds, 2 - Hearts, 3 - Spades
	Rank int // 2-14, 2-10, Jack, Queen, King, Ace
}
