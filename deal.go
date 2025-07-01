package main

import (
	// "bufio"
	"fmt"
	"math/rand"
	// "os"
)

type Cards []Card
type Hands []Cards
type Players []Player

type Deal struct {
	Hands       Hands
	Board       Cards
	HandResults []HandResult
}

type HandResult struct {
	Player           Player
	RelativeHandRank int
}

func getHandRankString(rank int) string {
	switch rank {
	case 1:
		return "straight flush"
	case 2:
		return "quads"
	case 3:
		return "full house"
	case 4:
		return "flush"
	case 5:
		return "straight"
	case 6:
		return "three of a kind"
	case 7:
		return "two pair"
	case 8:
		return "pair"
	case 9:
		return "high card"

	default:
		return "there's an error"
	}
}

func (d *Deck) ShuffleCards() {
	rand.Shuffle(len(d.Cards),
		func(i, j int) {
			d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
		})
}

// Flop
func (d *Deck) GetFlop() Cards {
	return d.BurnAndFlip(3)
}

// Turn
func (d *Deck) GetTurn() Cards {
	return d.BurnAndFlip(1)
}

// River
func (d *Deck) GetRiver() Cards {
	return d.BurnAndFlip(1)
}

// Reset CardIndex
func (d *Deck) ResetIndex() {
	d.CardIndex = 0
}

func (d *Deck) GetCard() Card {
	card := d.Cards[d.CardIndex]
	d.CardIndex++
	return card
}

func (d *Deck) BurnAndFlip(numCards int) Cards {

	var cards Cards

	//Burn a card:
	d.GetCard()

	for range numCards {
		nextCard := d.GetCard()
		cards = append(cards, nextCard)
	}
	return cards

}

func (d *Deck) DealCards(numHands, numCards int) Hands {
	var hands Hands
	for range numHands {
		hand := []Card{}
		hands = append(hands, hand)
	}

	for range numCards {
		for j := range numHands {
			currentHand := hands[j]
			nextCard := d.GetCard()
			currentHand = append(currentHand, nextCard)
			hands[j] = currentHand
		}
	}
	return hands
}

func (d *Deck) GetDeal(numPlayers int) Deal {

	hands := d.DealCards(numPlayers, 2)
	// fmt.Println("Press Enter to continue")
	// bufio.NewReader(os.Stdin).ReadBytes('\n')

	//wait turn
	flop := d.GetFlop()
	// fmt.Println("Press Enter to continue")
	// bufio.NewReader(os.Stdin).ReadBytes('\n')

	//wait turn
	turn := d.GetTurn()
	// fmt.Println("Press Enter to continue")
	// bufio.NewReader(os.Stdin).ReadBytes('\n')

	//wait turn
	river := d.GetRiver()
	// fmt.Println("Press Enter to continue")
	// bufio.NewReader(os.Stdin).ReadBytes('\n')

	//wait turn and end
	// fmt.Println("Press Enter to end game")
	// bufio.NewReader(os.Stdin).ReadBytes('\n')

	// Make turn based things
	// TODO - Google what every turn has

	var board Cards
	board = append(board, flop...)
	board = append(board, turn...)
	board = append(board, river...)

	var players Players

	for i, currentCards := range hands {
		var currentCardList Cards

		currentCardList = append(currentCardList, currentCards...)
		currentCardList = append(currentCardList, board...)

		bestFiveCardsOnTable, rank := GetFiveBestCards(currentCardList)

		currentPlayer := Player{
			Id:            i + 1,
			Name:          "PLACEHOLDER",
			BestFiveCards: bestFiveCardsOnTable,
			HandName:      getHandRankString(rank),
		}
		players = append(players, currentPlayer)
	}

	sortedPlayers := sortPlayers(players)

	winnerMap := getRankOrderMap(sortedPlayers)
	handResults := formatHandResults(winnerMap)

	deal := Deal{
		Hands:       hands,
		Board:       board,
		HandResults: handResults,
	}

	deal.PrintRanksAndBestFive()

	return deal
}

// sort players sorts the list of players to an ordered list based on ranking
func sortPlayers(pList Players) Players {

	playersList := make(Players, len(pList))
	copy(playersList, pList)

	for i := 0; i < len(playersList)-1; i++ {
		for j := 0; j < len(playersList)-i-1; j++ {

			curBestFive1 := playersList[j].BestFiveCards
			curBestFive2 := playersList[j+1].BestFiveCards

			winner, err := CompareTwoBestFive(curBestFive1, curBestFive2)
			if err != nil {
				fmt.Errorf("Error %v", err)
			}

			if winner == 2 {
				playersList[j], playersList[j+1] = playersList[j+1], playersList[j]
			}
		}
	}
	return playersList

}
func formatHandResults(p map[int]Players) []HandResult {
	numRanks := len(p)

	var handResults []HandResult

	for i := 0; i < numRanks; i++ {
		curRank := i + 1
		curPlayerList := p[curRank]

		for _, curPlayer := range curPlayerList {
			handResult := HandResult{
				Player:           curPlayer,
				RelativeHandRank: curRank,
			}
			handResults = append(handResults, handResult)
		}

	}
	return handResults

}

func getRankOrderMap(p Players) map[int]Players {
	var curPList Players
	curWinner := 1
	curPList = append(curPList, p[0])
	winnerMap := map[int]Players{curWinner: curPList}

	//start at the second element
	for i := 1; i < len(p); i++ {
		//keep track of cur rank being used
		curList := winnerMap[curWinner]
		//previous
		curBestFive1 := p[i-1].BestFiveCards
		//current
		curBestFive2 := p[i].BestFiveCards

		winner, err := CompareTwoBestFive(curBestFive1, curBestFive2)
		if err != nil {
			fmt.Errorf("Error %v", err)
		}

		if winner == 0 {
			curList = append(curList, p[i])
			winnerMap[curWinner] = curList
		} else {
			curWinner = curWinner + 1
			var newList Players
			newList = append(newList, p[i])
			winnerMap[curWinner] = newList
		}
	}
	return winnerMap
}
