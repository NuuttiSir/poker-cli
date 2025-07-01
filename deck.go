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

type ByNumber Cards

func (n ByNumber) Len() int { return len(n) }
func (n ByNumber) Less(i, j int) bool {

	var index1, index2 int

	for curIndex, val := range orderOfHighest {
		if val == n[i].Value {
			index1 = curIndex
		}

		if val == n[j].Value {
			index2 = curIndex
		}
	}

	return index1 < index2
}

func (n ByNumber) Swap(i, j int) { n[i], n[j] = n[j], n[i] }

// Sort highest to lowest based on Value param
func (c *Cards) Sort() {

}
