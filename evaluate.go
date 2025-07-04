package main

import (
	"fmt"
	"sort"
	"strconv"
)

var orderOfHighest []string

func init() {
	//Ace represented twice - A for high ace and 1 for low ace
	orderOfHighest = []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2", "1"}

}

// GetFiveBest evaluates the hand and prints out what it is
// First return param: the 5 best cards
// Second return param: the ranking of the 5 best cards.  Rankings can be found above
func GetFiveBestCards(cards Cards) (Cards, int) {

	cardCopy := make(Cards, len(cards))
	copy(cardCopy, cards)

	sort.Sort(ByNumber(cardCopy))

	//Get suit map - creates four subsets of all the cards for eval
	suitMap := getSuitToCardsMap(cardCopy)

	//Check straight flush
	isStraightFlush, straightFlushCards := checkForStraightFlush(suitMap)
	if isStraightFlush {
		return straightFlushCards, straightFlush
	}

	//Check quads
	isQuads, cardsFound := checkHighestCardForQuantity(cardCopy, 4)

	if isQuads {

		cardCopy.Remove(cardsFound)
		highCards := getNumHighCards(cardCopy, 1)
		cardCopy.Add(cardsFound)
		cardsFound = append(cardsFound, highCards...)

		return cardsFound, quads
	}

	//Check full house
	isThreeOfAKind, foundCards := checkHighestCardForQuantity(cardCopy, 3)
	if isThreeOfAKind {

		cardCopy.Remove(foundCards)

		isPair, foundPair := checkHighestCardForQuantity(cardCopy, 2)

		//if not a full house, add back 3 of a kind for later eval
		cardCopy.Add(foundCards)

		if isPair {
			foundCards = append(foundCards, foundPair...)
			return foundCards, fullHouse
		}

	}

	//Check flush
	isFlush, flushCards := checkForFlush(suitMap)
	if isFlush {
		return flushCards, flush
	}

	//=======================================================================
	// below this line, suit no longer matters (but would like to return winning hand with suit)

	//Check straight
	isStraight, straightCards := checkForFiveInARow(cardCopy)
	if isStraight {
		return straightCards, straight
	}

	isThreeOfAKind, foundCards = checkHighestCardForQuantity(cardCopy, 3)
	if isThreeOfAKind {

		cardCopy.Remove(foundCards)
		twoHighCards := getNumHighCards(cardCopy, 2)
		cardCopy.Add(foundCards)
		foundCards = append(foundCards, twoHighCards...)
		return foundCards, threeOfAKind

	}

	//Processes Pair and Two Pair
	isPair, foundCards := checkHighestCardForQuantity(cardCopy, 2)
	if isPair {

		cardCopy.Remove(foundCards)
		isTwoPair, secondPair := checkHighestCardForQuantity(cardCopy, 2)

		if isTwoPair {

			cardCopy.Remove(secondPair)
			foundCards = append(foundCards, secondPair...)

			highCards := getNumHighCards(cardCopy, 1)
			cardCopy.Remove(highCards)
			foundCards = append(foundCards, highCards...)

			cardCopy.Add(foundCards)
			return foundCards, twoPair
		}

		//Pair (not two pair)
		highCards := getNumHighCards(cardCopy, 3)

		cardCopy.Remove(highCards)
		foundCards = append(foundCards, highCards...)

		cardCopy.Add(foundCards)
		return foundCards, pair

	}

	//just a high card
	highCards := getNumHighCards(cardCopy, 5)
	return highCards, highCard
}

// Add adds cards to a Cards object.  This is intended to be used for failed multistep checks (fullhouse, two pair)
func (c *Cards) Add(cardsToAdd Cards) {

	for _, cardToAdd := range cardsToAdd {
		(*c) = append((*c), cardToAdd)
	}

}

// Remove removes cards from a Cards object
func (c *Cards) Remove(cardsToRemove Cards) {

	for _, cardToRemove := range cardsToRemove {

		for i, curCard := range *c {

			if cardToRemove.Suit == curCard.Suit &&
				cardToRemove.Value == curCard.Value {
				(*c)[i] = (*c)[len(*c)-1]
				(*c)[len(*c)-1] = Card{}
				*c = (*c)[:len(*c)-1]
			}

		}
	}
}

func checkForFlush(suitMap map[string]Cards) (bool, Cards) {
	for _, value := range suitMap {
		if len(value) >= 5 {
			sort.Sort(ByNumber(value))
			return true, value[:5]
		}
	}
	return false, Cards{}
}

func getNumHighCards(cards Cards, highCardsNeeded int) Cards {

	var foundCards Cards

	for range highCardsNeeded {
		foundHighCard, highCard := checkHighestCardForQuantity(cards, 1)

		if foundHighCard {
			foundCards = append(foundCards, highCard[0])
		} else {
			fmt.Println("oof somethings broken")
		}
		cards.Remove(highCard)

	}

	return foundCards
}

func checkHighestCardForQuantity(cards Cards, cardsNeeded int) (bool, Cards) {

	numMap := cards.getCardValues()

	var highCard string
	for _, value := range orderOfHighest {

		if numMap[value] >= cardsNeeded {
			highCard = value
			break
		}
	}

	if highCard == "" {
		return false, Cards{}
	}

	var foundCards Cards

	for _, curCard := range cards {
		if curCard.Value == highCard {

			foundCards = append(foundCards, curCard)

			if len(foundCards) == cardsNeeded {
				break
			}

		}
	}

	return true, foundCards
}

func getSuitToCardsMap(cards Cards) map[string]Cards {

	var suitMap map[string]Cards

	suitMap = make(map[string]Cards)

	suitMap["H"] = Cards{}
	suitMap["S"] = Cards{}
	suitMap["C"] = Cards{}
	suitMap["D"] = Cards{}

	for _, card := range cards {

		curList := suitMap[card.Suit]
		curList = append(curList, card)

		suitMap[card.Suit] = curList
	}
	return suitMap
}

func checkForStraightFlush(suitMap map[string]Cards) (bool, Cards) {

	for _, cards := range suitMap {

		//sort the cards of the same suit
		sort.Sort(ByNumber(cards))

		if len(cards) > 0 {
			straightFlushFound, straightFlushCards := checkForFiveInARow(cards)
			if straightFlushFound {
				return straightFlushFound, straightFlushCards
			}
		}
	}
	return false, Cards{}
}

func (c Cards) getCardValues() map[string]int {

	numMap := map[string]int{
		"A": 0,
		"K": 0,
		"Q": 0,
		"J": 0,
		"T": 0,
		"9": 0,
		"8": 0,
		"7": 0,
		"6": 0,
		"5": 0,
		"4": 0,
		"3": 0,
		"2": 0,
	}

	for _, card := range c {
		numMap[card.Value]++
	}
	return numMap
}

// GetNumberValues returns list of ints for
func (c Cards) getNumberValues() ([]int, error) {

	var values []int

	for _, card := range c {
		var curVal int

		curVal = getNumberValue(card)
		values = append(values, curVal)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	return values, nil
}

func getNumberValue(card Card) int {
	var curVal int
	var err error
	switch card.Value {
	case "T":
		curVal = 10
	case "J":
		curVal = 11
	case "Q":
		curVal = 12
	case "K":
		curVal = 13
	case "A":
		curVal = 14
	default:
		curVal, err = strconv.Atoi(card.Value)
		if err != nil {
			fmt.Println("error converting card to number value: " + card.Value + card.Suit + ".  Error: " + err.Error())
		}
	}
	return curVal
}

// if 5 in a row, returns true and the number
// if not, returns false and 0
func checkForFiveInARow(cards Cards) (bool, Cards) {

	var fiveInARow Cards
	sort.Sort(ByNumber(cards))
	//add value for 1 for an ace (to check for low straight - ace can be high or low)
	if cards[0].Value == "A" {
		newCard := Card{
			Suit:  cards[0].Suit,
			Value: "1",
		}
		cards = append(cards, newCard)
	}

	fiveInARow = append(fiveInARow, cards[0])

	for i := 0; i < len(cards)-1; i++ {
		card1 := getNumberValue(cards[i])
		card2 := getNumberValue(cards[i+1])

		//if sequential values are 1 apart, add card to array
		if card1-1 == card2 {
			fiveInARow = append(fiveInARow, cards[i+1])

			//if values are the same, don't add but also dont reset
		} else if card1 == card2 {
			continue
		} else {
			//reset to next value
			fiveInARow = fiveInARow[:0]
			fiveInARow = append(fiveInARow, cards[i+1])
		}
		if len(fiveInARow) == 5 {

			return true, fiveInARow
		}
	}

	return false, Cards{}
}
