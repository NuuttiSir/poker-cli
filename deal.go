package main

import (
	"fmt"
	"math/rand/v2"
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
	//wait turn
	flop := d.GetFlop()
	//wait turn
	turn := d.GetTurn()
	//wait turn
	river := d.GetRiver()
	//wait turn and end

	var board Cards
	board = append(board, flop...)
	board = append(board, turn...)
	board = append(board, river...)

	var players Players

	for i, currentCards := range hands {
		var currentCardList Cards

		currentCardList = append(currentCardList, currentCards...)
		currentCardList = append(currentCardList, board...)

		//GEt best hand

		currentPlayer := Player{
			Id: i + 1,
			// BEstFIveCardsInHand
			HandName: "", // GetStringForHandRank
		}
		players = append(players, currentPlayer)
	}

	// SortPlayerByHandRank
	// GetWinnerHand and show

	deal := Deal{
		Hands:       hands,
		Board:       board,
		HandResults: nil, // Get hand results ?
	}
	//Printing for the hell of it REMOVE JOSKUS
	fmt.Println(hands)
	fmt.Println(board)
	fmt.Println(players)
	fmt.Println(deal)

	return deal
}
