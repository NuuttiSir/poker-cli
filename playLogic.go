package main

func Play() {
	var deck Deck

	deck = CreateDeck()
	deck.ShuffleCards()
	deck.GetDeal(2)

	Menu()
}
