package main

import (
	"errors"
	"fmt"
	"sort"
)

const (
	straightFlush = 1
	quads         = 2
	fullHouse     = 3
	flush         = 4
	straight      = 5
	threeOfAKind  = 6
	twoPair       = 7
	pair          = 8
	highCard      = 9
	errConstant   = -1
)

// CompareTwoBestFive compares two hands
// returns 1 if first hand best, 2 if second hand best and
// 0 if hands are the same (in evaluation, not necessarily identical)
// errConstant if error
func CompareTwoBestFive(firstFive, secondFive Cards) (int, error) {

	firstFiveCopy := make(Cards, len(firstFive))
	copy(firstFiveCopy, firstFive)
	secondFiveCopy := make(Cards, len(secondFive))
	copy(secondFiveCopy, secondFive)

	if len(firstFiveCopy) != 5 || len(secondFiveCopy) != 5 {
		return errConstant, errors.New("bad input - both card sets need to be of length 5")
	}

	firstFiveCopy, rank1 := GetFiveBestCards(firstFiveCopy)
	secondFiveCopy, rank2 := GetFiveBestCards(secondFiveCopy)
	if rank1 < rank2 {
		return 1, nil
	} else if rank2 < rank1 {
		return 2, nil
	}

	//have same level of hand, need to evaluate further
	switch rank1 {
	case straightFlush:
		return compareStraightFlushes(firstFiveCopy, secondFiveCopy), nil
	case quads:
		return compareQuads(firstFiveCopy, secondFiveCopy), nil
	case fullHouse:
		return compareFullHouses(firstFiveCopy, secondFiveCopy), nil
	case flush:
		return compareFlushes(firstFiveCopy, secondFiveCopy), nil
	case straight:
		return compareStraight(firstFiveCopy, secondFiveCopy), nil
	case threeOfAKind:
		return compareThreeOfAKind(firstFiveCopy, secondFiveCopy), nil
	case twoPair:
		return compareTwoPair(firstFiveCopy, secondFiveCopy), nil
	case pair:
		return comparePair(firstFiveCopy, secondFiveCopy), nil
	case highCard:
		return compareHighCards(firstFiveCopy, secondFiveCopy), nil
	}

	return errConstant, errors.New("do not understand input")

}

func compareStraightFlushes(firstFive Cards, secondFive Cards) int {
	return compareStraight(firstFive, secondFive)
}

func compareStraight(firstFive Cards, secondFive Cards) int {
	sort.Sort(ByNumber(firstFive))

	sort.Sort(ByNumber(secondFive))

	//using index 1 because its the simple fix for the low straight case
	//the sort will treat ace as high
	return compareCard(firstFive[1], secondFive[1])
}

func compareQuads(firstFive Cards, secondFive Cards) int {
	quads1, quads2, err := getHighestCardsForQuantity(firstFive, secondFive, 4)
	if err != nil {
		return errConstant
	}

	//remove the quads from both hands and get the high card remaining value
	firstFive.Remove(quads1)
	secondFive.Remove(quads2)

	highCard1, highCard2, err := getHighestCardsForQuantity(firstFive, secondFive, 1)
	if err != nil {
		return errConstant
	}
	//Add back quads
	firstFive.Add(quads1)
	secondFive.Add(quads2)

	//Make list of cards to compare
	var firstEvalOrder, secondEvalOrder Cards
	firstEvalOrder = append(firstEvalOrder, quads1[0], highCard1[0])
	secondEvalOrder = append(secondEvalOrder, quads2[0], highCard2[0])

	return compareCards(firstEvalOrder, secondEvalOrder)
}

func compareFullHouses(firstFive, secondFive Cards) int {
	threeOfAKind1, threeOfAKind2, err := getHighestCardsForQuantity(firstFive, secondFive, 3)
	if err != nil {
		return errConstant
	}

	//remove for pair eval
	firstFive.Remove(threeOfAKind1)
	secondFive.Remove(threeOfAKind2)

	pair1, pair2, err := getHighestCardsForQuantity(firstFive, secondFive, 2)
	if err != nil {
		return errConstant
	}

	//add back
	firstFive.Add(threeOfAKind1)
	secondFive.Add(threeOfAKind2)

	//Make list of cards to compare
	var firstEvalOrder, secondEvalOrder Cards
	firstEvalOrder = append(firstEvalOrder, threeOfAKind1[0], pair1[0])
	secondEvalOrder = append(secondEvalOrder, threeOfAKind2[0], pair2[0])

	return compareCards(firstEvalOrder, secondEvalOrder)
}

func compareFlushes(firstFive, secondFive Cards) int {

	//Sort from highest to lowest
	sort.Sort(ByNumber(firstFive))
	sort.Sort(ByNumber(secondFive))

	for i := range firstFive {

		//Go in order and compare each value
		result := compareCard((firstFive)[i], (secondFive)[i])
		if result != 0 {
			return result
		}
	}
	return 0
}

func compareThreeOfAKind(firstFive, secondFive Cards) int {
	threeOfAKind1, threeOfAKind2, err := getHighestCardsForQuantity(firstFive, secondFive, 3)
	if err != nil {
		return errConstant
	}

	result := compareCard(threeOfAKind1[0], threeOfAKind2[0])
	if result != 0 {
		return result
	}

	//remove three of a kind for high card eval
	firstFive.Remove(threeOfAKind1)
	secondFive.Remove(threeOfAKind2)

	resultHighCards := compareHighCards(firstFive, secondFive)
	firstFive.Add(threeOfAKind1)
	secondFive.Add(threeOfAKind2)
	return resultHighCards

}

// Eval criteria: first pair comparison, then second pair comparison, then high card comparison
func compareTwoPair(firstFive, secondFive Cards) int {
	//first pair
	firstPair1, firstPair2, err := getHighestCardsForQuantity(firstFive, secondFive, 2)
	if err != nil {
		return errConstant
	}
	firstFive.Remove(firstPair1)
	secondFive.Remove(firstPair2)

	//then second pair
	secondPair1, secondPair2, err := getHighestCardsForQuantity(firstFive, secondFive, 2)
	if err != nil {
		return errConstant
	}
	firstFive.Remove(secondPair1)
	secondFive.Remove(secondPair2)

	//then high card
	highCard1, highCard2, err := getHighestCardsForQuantity(firstFive, secondFive, 1)
	if err != nil {
		return errConstant
	}

	//Add cards back that were removed
	firstFive.Add(firstPair1)
	secondFive.Add(firstPair2)
	firstFive.Add(secondPair1)
	secondFive.Add(secondPair2)

	//Make list of cards to compare
	var firstEvalOrder, secondEvalOrder Cards
	firstEvalOrder = append(firstEvalOrder, firstPair1[0], secondPair1[0], highCard1[0])
	secondEvalOrder = append(secondEvalOrder, firstPair2[0], secondPair2[0], highCard2[0])

	return compareCards(firstEvalOrder, secondEvalOrder)
}

// Eval criteria: first pair comparison, then second pair comparison, then high card comparison
func comparePair(firstFive, secondFive Cards) int {
	//pair
	pair1, pair2, err := getHighestCardsForQuantity(firstFive, secondFive, 2)
	if err != nil {
		return errConstant
	}
	firstFive.Remove(pair1)
	secondFive.Remove(pair2)

	result := compareCard(pair1[0], pair2[0])
	if result != 0 {
		//add back removed cards
		firstFive.Add(pair1)
		secondFive.Add(pair2)
		return result
	}

	result1 := compareHighCards(firstFive, secondFive)
	//add back removed cards
	firstFive.Add(pair1)
	secondFive.Add(pair2)
	return result1
}

func compareHighCards(firstFive, secondFive Cards) int {

	sort.Sort(ByNumber(firstFive))
	sort.Sort(ByNumber(secondFive))

	return compareCards(firstFive, secondFive)
}

//compareCards takes two lists of equal length cards and compares each index

//expects cards to be in order from first to eval to last (not necessarily in order of highest)
//Example: this hand:
//KC, TD, AS, 4H, 4D, should be passed in as 4H, 4D, AS, KC, TD
//so that it can first evaluate the pair then the high cards

// this function is only to be used by cards list of same length
func compareCards(cardList1, cardList2 Cards) int {
	for index := range cardList1 {
		result := compareCard(cardList1[index], cardList2[index])
		if result != 0 {
			return result
		}
	}

	return 0
}

func getHighestCardsForQuantity(firstCards, secondCards Cards, quantity int) (Cards, Cards, error) {

	found1, cards1 := checkHighestCardForQuantity(firstCards, quantity)
	found2, cards2 := checkHighestCardForQuantity(secondCards, quantity)
	if !found1 || !found2 {
		err := errors.New("bad input - function can only be called if its known to have the desired quantity")
		fmt.Println("error: " + err.Error())
		return Cards{}, Cards{}, err
	}

	if len(cards1) == 0 || len(cards2) == 0 {
		err := errors.New("bad input - no cards found")
		fmt.Println("error: " + err.Error())
		return Cards{}, Cards{}, err
	}

	return cards1, cards2, nil
}

// compareCard iterates through the orderOfHighest list
// returns 1 if first card earlier in this list, 2 if second card earlier in the list
// and 0 if they appear in the same spot as this list
func compareCard(card1 Card, card2 Card) int {

	//orderOfHighest = []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2", "1"}

	var index1, index2 int

	for i, val := range orderOfHighest {
		if val == card1.Value {
			index1 = i
		}
		if val == card2.Value {
			index2 = i
		}
	}

	if index1 < index2 {
		return 1
	} else if index2 < index1 {
		return 2
	} else {
		return 0
	}

}
