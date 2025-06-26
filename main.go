package main

import (
	"fmt"
	"math/rand"
)

type Hands []Cards

type Cards []Card

type Players []Player

type Card struct {
	Suit  string
	Value string
}

type Deck struct {
	Cards     []Card
	CardIndex int
}

type Player struct {
	Id       int
	HandName string
}

type Deal struct {
	Hands       Hands
	Board       Cards
	HandResults []HandResult
}

type HandResult struct {
	Player           Player
	RelativeHandRank int
}

func Menu() {
	fmt.Println("Welcome to my poker cli game")
	fmt.Println()
	fmt.Println("Press the corresponding number to select")
	fmt.Println("1. Play\n2. Help\n3. Quitn")
}

func HelpMenu() {
	fmt.Println("Here are the rules of poker and hand rankings etc")
}

func Quit() {
	fmt.Println("Thank you for playing")
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

func (d *Deck) ShuffleCards() {
	rand.Shuffle(len(d.Cards),
		func(i, j int) {
			d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
		})
}

func (d *Deck) GetCard() Card {
	card := d.Cards[d.CardIndex]
	d.CardIndex++
	return card
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
	fmt.Println(hands)
	return hands
}

func main() {
	var deck Deck
	deck = CreateDeck()
	deck.ShuffleCards()
	deck.DealCards(2, 2)
}
