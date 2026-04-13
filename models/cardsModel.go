package models

type Suit int

const (
	heart Suit = iota
	diamond
	spade
	club
	trump
)

type Card struct {
	ID      int
	Name    string
	Suit    Suit
	IsTrump bool
	IsBout  bool
	Power   int
	Points  float32
}

// A function so services can receive the cards
func NewDeck() []Card {
	var cards = []Card{
		// Trumps (Atouts)
		{ID: 0, Name: "Excuse", Suit: trump, IsTrump: true, IsBout: true, Power: 0, Points: 4.5},
		{ID: 1, Name: "Petit", Suit: trump, IsTrump: true, IsBout: true, Power: 57, Points: 4.5},
		{ID: 2, Name: "Trump 2", Suit: trump, IsTrump: true, IsBout: false, Power: 58, Points: 0.5},
		{ID: 3, Name: "Trump 3", Suit: trump, IsTrump: true, IsBout: false, Power: 59, Points: 0.5},
		{ID: 4, Name: "Trump 4", Suit: trump, IsTrump: true, IsBout: false, Power: 60, Points: 0.5},
		{ID: 5, Name: "Trump 5", Suit: trump, IsTrump: true, IsBout: false, Power: 61, Points: 0.5},
		{ID: 6, Name: "Trump 6", Suit: trump, IsTrump: true, IsBout: false, Power: 62, Points: 0.5},
		{ID: 7, Name: "Trump 7", Suit: trump, IsTrump: true, IsBout: false, Power: 63, Points: 0.5},
		{ID: 8, Name: "Trump 8", Suit: trump, IsTrump: true, IsBout: false, Power: 64, Points: 0.5},
		{ID: 9, Name: "Trump 9", Suit: trump, IsTrump: true, IsBout: false, Power: 65, Points: 0.5},
		{ID: 10, Name: "Trump 10", Suit: trump, IsTrump: true, IsBout: false, Power: 66, Points: 0.5},
		{ID: 11, Name: "Trump 11", Suit: trump, IsTrump: true, IsBout: false, Power: 67, Points: 0.5},
		{ID: 12, Name: "Trump 12", Suit: trump, IsTrump: true, IsBout: false, Power: 68, Points: 0.5},
		{ID: 13, Name: "Trump 13", Suit: trump, IsTrump: true, IsBout: false, Power: 69, Points: 0.5},
		{ID: 14, Name: "Trump 14", Suit: trump, IsTrump: true, IsBout: false, Power: 70, Points: 0.5},
		{ID: 15, Name: "Trump 15", Suit: trump, IsTrump: true, IsBout: false, Power: 71, Points: 0.5},
		{ID: 16, Name: "Trump 16", Suit: trump, IsTrump: true, IsBout: false, Power: 72, Points: 0.5},
		{ID: 17, Name: "Trump 17", Suit: trump, IsTrump: true, IsBout: false, Power: 73, Points: 0.5},
		{ID: 18, Name: "Trump 18", Suit: trump, IsTrump: true, IsBout: false, Power: 74, Points: 0.5},
		{ID: 19, Name: "Trump 19", Suit: trump, IsTrump: true, IsBout: false, Power: 75, Points: 0.5},
		{ID: 20, Name: "Trump 20", Suit: trump, IsTrump: true, IsBout: false, Power: 76, Points: 0.5},
		{ID: 21, Name: "Trump 21", Suit: trump, IsTrump: true, IsBout: true, Power: 77, Points: 4.5},

		// Hearts (Coeurs)
		{ID: 22, Name: "1 of Hearts", Suit: heart, IsTrump: false, IsBout: false, Power: 1, Points: 0.5},
		{ID: 23, Name: "2 of Hearts", Suit: heart, IsTrump: false, IsBout: false, Power: 2, Points: 0.5},
		{ID: 24, Name: "3 of Hearts", Suit: heart, IsTrump: false, IsBout: false, Power: 3, Points: 0.5},
		{ID: 25, Name: "4 of Hearts", Suit: heart, IsTrump: false, IsBout: false, Power: 4, Points: 0.5},
		{ID: 26, Name: "5 of Hearts", Suit: heart, IsTrump: false, IsBout: false, Power: 5, Points: 0.5},
		{ID: 27, Name: "6 of Hearts", Suit: heart, IsTrump: false, IsBout: false, Power: 6, Points: 0.5},
		{ID: 28, Name: "7 of Hearts", Suit: heart, IsTrump: false, IsBout: false, Power: 7, Points: 0.5},
		{ID: 29, Name: "8 of Hearts", Suit: heart, IsTrump: false, IsBout: false, Power: 8, Points: 0.5},
		{ID: 30, Name: "9 of Hearts", Suit: heart, IsTrump: false, IsBout: false, Power: 9, Points: 0.5},
		{ID: 31, Name: "10 of Hearts", Suit: heart, IsTrump: false, IsBout: false, Power: 10, Points: 0.5},
		{ID: 32, Name: "Jack of Hearts", Suit: heart, IsTrump: false, IsBout: false, Power: 11, Points: 1.5},
		{ID: 33, Name: "Knight of Hearts", Suit: heart, IsTrump: false, IsBout: false, Power: 12, Points: 2.5},
		{ID: 34, Name: "Queen of Hearts", Suit: heart, IsTrump: false, IsBout: false, Power: 13, Points: 3.5},
		{ID: 35, Name: "King of Hearts", Suit: heart, IsTrump: false, IsBout: false, Power: 14, Points: 4.5},

		// Diamonds (Carreaux)
		{ID: 36, Name: "1 of Diamonds", Suit: diamond, IsTrump: false, IsBout: false, Power: 1, Points: 0.5},
		{ID: 37, Name: "2 of Diamonds", Suit: diamond, IsTrump: false, IsBout: false, Power: 2, Points: 0.5},
		{ID: 38, Name: "3 of Diamonds", Suit: diamond, IsTrump: false, IsBout: false, Power: 3, Points: 0.5},
		{ID: 39, Name: "4 of Diamonds", Suit: diamond, IsTrump: false, IsBout: false, Power: 4, Points: 0.5},
		{ID: 40, Name: "5 of Diamonds", Suit: diamond, IsTrump: false, IsBout: false, Power: 5, Points: 0.5},
		{ID: 41, Name: "6 of Diamonds", Suit: diamond, IsTrump: false, IsBout: false, Power: 6, Points: 0.5},
		{ID: 42, Name: "7 of Diamonds", Suit: diamond, IsTrump: false, IsBout: false, Power: 7, Points: 0.5},
		{ID: 43, Name: "8 of Diamonds", Suit: diamond, IsTrump: false, IsBout: false, Power: 8, Points: 0.5},
		{ID: 44, Name: "9 of Diamonds", Suit: diamond, IsTrump: false, IsBout: false, Power: 9, Points: 0.5},
		{ID: 45, Name: "10 of Diamonds", Suit: diamond, IsTrump: false, IsBout: false, Power: 10, Points: 0.5},
		{ID: 46, Name: "Jack of Diamonds", Suit: diamond, IsTrump: false, IsBout: false, Power: 11, Points: 1.5},
		{ID: 47, Name: "Knight of Diamonds", Suit: diamond, IsTrump: false, IsBout: false, Power: 12, Points: 2.5},
		{ID: 48, Name: "Queen of Diamonds", Suit: diamond, IsTrump: false, IsBout: false, Power: 13, Points: 3.5},
		{ID: 49, Name: "King of Diamonds", Suit: diamond, IsTrump: false, IsBout: false, Power: 14, Points: 4.5},

		// Spades (Pics)
		{ID: 50, Name: "1 of Spades", Suit: spade, IsTrump: false, IsBout: false, Power: 1, Points: 0.5},
		{ID: 51, Name: "2 of Spades", Suit: spade, IsTrump: false, IsBout: false, Power: 2, Points: 0.5},
		{ID: 52, Name: "3 of Spades", Suit: spade, IsTrump: false, IsBout: false, Power: 3, Points: 0.5},
		{ID: 53, Name: "4 of Spades", Suit: spade, IsTrump: false, IsBout: false, Power: 4, Points: 0.5},
		{ID: 54, Name: "5 of Spades", Suit: spade, IsTrump: false, IsBout: false, Power: 5, Points: 0.5},
		{ID: 55, Name: "6 of Spades", Suit: spade, IsTrump: false, IsBout: false, Power: 6, Points: 0.5},
		{ID: 56, Name: "7 of Spades", Suit: spade, IsTrump: false, IsBout: false, Power: 7, Points: 0.5},
		{ID: 57, Name: "8 of Spades", Suit: spade, IsTrump: false, IsBout: false, Power: 8, Points: 0.5},
		{ID: 58, Name: "9 of Spades", Suit: spade, IsTrump: false, IsBout: false, Power: 9, Points: 0.5},
		{ID: 59, Name: "10 of Spades", Suit: spade, IsTrump: false, IsBout: false, Power: 10, Points: 0.5},
		{ID: 60, Name: "Jack of Spades", Suit: spade, IsTrump: false, IsBout: false, Power: 11, Points: 1.5},
		{ID: 61, Name: "Knight of Spades", Suit: spade, IsTrump: false, IsBout: false, Power: 12, Points: 2.5},
		{ID: 62, Name: "Queen of Spades", Suit: spade, IsTrump: false, IsBout: false, Power: 13, Points: 3.5},
		{ID: 63, Name: "King of Spades", Suit: spade, IsTrump: false, IsBout: false, Power: 14, Points: 4.5},

		// Clubs (Trèfles)
		{ID: 64, Name: "1 of Clubs", Suit: club, IsTrump: false, IsBout: false, Power: 1, Points: 0.5},
		{ID: 65, Name: "2 of Clubs", Suit: club, IsTrump: false, IsBout: false, Power: 2, Points: 0.5},
		{ID: 66, Name: "3 of Clubs", Suit: club, IsTrump: false, IsBout: false, Power: 3, Points: 0.5},
		{ID: 67, Name: "4 of Clubs", Suit: club, IsTrump: false, IsBout: false, Power: 4, Points: 0.5},
		{ID: 68, Name: "5 of Clubs", Suit: club, IsTrump: false, IsBout: false, Power: 5, Points: 0.5},
		{ID: 69, Name: "6 of Clubs", Suit: club, IsTrump: false, IsBout: false, Power: 6, Points: 0.5},
		{ID: 70, Name: "7 of Clubs", Suit: club, IsTrump: false, IsBout: false, Power: 7, Points: 0.5},
		{ID: 71, Name: "8 of Clubs", Suit: club, IsTrump: false, IsBout: false, Power: 8, Points: 0.5},
		{ID: 72, Name: "9 of Clubs", Suit: club, IsTrump: false, IsBout: false, Power: 9, Points: 0.5},
		{ID: 73, Name: "10 of Clubs", Suit: club, IsTrump: false, IsBout: false, Power: 10, Points: 0.5},
		{ID: 74, Name: "Jack of Clubs", Suit: club, IsTrump: false, IsBout: false, Power: 11, Points: 1.5},
		{ID: 75, Name: "Knight of Clubs", Suit: club, IsTrump: false, IsBout: false, Power: 12, Points: 2.5},
		{ID: 76, Name: "Queen of Clubs", Suit: club, IsTrump: false, IsBout: false, Power: 13, Points: 3.5},
		{ID: 77, Name: "King of Clubs", Suit: club, IsTrump: false, IsBout: false, Power: 14, Points: 4.5},
	}
	return cards
}
