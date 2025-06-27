package poker

import (
	"fmt"

	"github.com/NuuttiSir/poker-cli/UI"
)

func Play() {
	var input int
	var deck Deck

	deck = CreateDeck()
	deck.ShuffleCards()
	deck.GetDeal(2)

	UI.Menu()
	fmt.Scan(&input)

	switch input {
	case 1:
		Play()
	case 2:
		UI.HelpMenu()
	case 3:
		fmt.Println("Thanks for playing")
		UI.Quit()
	}
}
