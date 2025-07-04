package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

var (
	MIN_SPACE = 16
)

// PrintOrder prints the order of the deck
func (d Deck) PrintOrder() {
	var b bytes.Buffer
	b.WriteString("Card Order:\n")
	for index, card := range d.Cards {
		b.WriteString(card.Value)
		b.WriteString(card.Suit)

		//b.WriteString("\n")
		if (index+1)%13 == 0 {
			b.WriteString("\n")
		} else {
			b.WriteString(", ")
		}
	}
	fmt.Println(b.String())
}

// PrintRemainingCards prints the remaining cards in the deck
func (d Deck) PrintRemainingCards() {
	var b bytes.Buffer
	b.WriteString("Remaining Cards:\n")
	remainingCards := 64 - d.CardIndex

	mod := remainingCards / 6
	for index, card := range d.Cards {
		if index < d.CardIndex {
			continue
		}
		b.WriteString(card.Value)
		b.WriteString(card.Suit)

		//make the next card the "0" placement
		if (index+1-d.CardIndex)%mod == 0 {
			b.WriteString("\n")
		} else {
			b.WriteString(", ")
		}
	}
	fmt.Println(b.String())
}

// Print prints hands
func (h Hands) Print() {
	for index, hand := range h {
		playerStr := strconv.Itoa(index + 1)

		//Add space to make formatting prettier if 10 handed - 2 chars vs 1 in digit
		if len(playerStr) == 1 {
			playerStr = " " + playerStr
		}
		hand.Print("Hand "+playerStr, "")
	}
	fmt.Printf("\n")
}

// Print prints cards
func (c Cards) Print(beforeStr, afterStr string) {
	var b bytes.Buffer
	b.WriteString(beforeStr)
	b.WriteString(": ")
	for index, card := range c {
		if card.Value == "1" {
			b.WriteString("A")
		} else {
			b.WriteString(card.Value)
		}
		b.WriteString(card.Suit)
		if index != len(c)-1 {
			b.WriteString(", ")
		}

	}
	b.WriteString(afterStr)
	fmt.Println(b.String())

}

func (c Cards) getBestFiveString() string {
	var b bytes.Buffer
	for index, card := range c {
		if card.Value == "1" {
			b.WriteString("A")
		} else {
			b.WriteString(card.Value)
		}
		b.WriteString(card.Suit)
		if index != len(c)-1 {
			b.WriteString(" ")
		}

	}
	return b.String()
}

// Print prints players
func (p Players) Print() {
	for i := range p {
		playerNum := p[i].Id
		p[i].BestFiveCards.Print("Player "+strconv.Itoa(playerNum), " ("+p[i].HandName+")")
	}
}

// PrintBoard prints the board for a game
func (d *Deal) PrintBoard() {
	d.Board.Print("Board", "")
}

// PrintHands prints the hands for a game
func (d *Deal) PrintHands() {
	d.Hands.Print()
}

func (d *Deal) PrintRanksAndBestFive() {
	fmt.Println("Results:")
	fmt.Println("====================================================")
	fmt.Println("| Player | Rank |    Best Five   |     Hand Name   |")
	fmt.Println("====================================================")

	for _, handResult := range d.HandResults {
		line := getLineToPrint(handResult)
		fmt.Println(line)
	}
	fmt.Println("----------------------------------------------------")
}

func getLineToPrint(handResult HandResult) string {
	curPlayer := handResult.Player

	//Add space to make formatting prettier if 10 handed - 2 chars vs 1 in digit
	playerNumStr := strconv.Itoa(curPlayer.Id)
	if len(playerNumStr) == 1 {
		playerNumStr = playerNumStr + " "
	}

	//Add space to make formatting prettier if 10 handed - 2 chars vs 1 in digit
	handRankStr := strconv.Itoa(handResult.RelativeHandRank)
	if len(handRankStr) == 1 {
		handRankStr = " " + handRankStr
	}

	bestFiveStr := curPlayer.BestFiveCards.getBestFiveString()

	//Calculate spaces before and after hand name for prettier formatting
	numSpaces := MIN_SPACE - len(curPlayer.HandName)
	spacesBefore := strings.Repeat(" ", numSpaces/2)
	var spacesAfter string
	if numSpaces%2 == 0 {
		spacesAfter = strings.Repeat(" ", numSpaces/2)
	} else {
		spacesAfter = strings.Repeat(" ", numSpaces/2+1)

	}

	//Build output string and return
	var sb strings.Builder
	sb.WriteString("|   ")
	sb.WriteString(playerNumStr)
	sb.WriteString("   |  ")
	sb.WriteString(handRankStr)
	sb.WriteString("  | ")
	sb.WriteString(bestFiveStr)
	sb.WriteString(" | ")
	sb.WriteString(spacesBefore)
	sb.WriteString(curPlayer.HandName)
	sb.WriteString(spacesAfter)
	sb.WriteString("|")
	return sb.String()
}

// PrintBoardAndHands prints the board and the hands
func (d *Deal) PrintBoardAndHands() {
	d.PrintBoard()
	d.PrintHands()
}
