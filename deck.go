package main

type Card struct {
	Suit  string
	Value string
}

type Deck struct {
	Cards     []Card
	CardIndex int
}

func CreateDeck() Deck {
	suits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	var deck Deck

	for _, suit := range suits {
		for _, value := range values {
			currentCard := Card{
				Suit:  suit,
				Value: value,
			}
			deck.Cards = append(deck.Cards, currentCard)
		}
	}
	deck.CardIndex = 0
	return deck
}
