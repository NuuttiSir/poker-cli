package main

import (
	"fmt"
)

func Menu() {
	var input int

	fmt.Println("Welcome to my poker cli game")
	fmt.Println()
	fmt.Println("Press the corresponding number to select")
	fmt.Println("1. Play\n2. Help\n3. Quit")

	fmt.Scan(&input)

	switch input {
	case 1:
		Play()
	case 2:
		HelpMenu()
	case 3:
		Quit()
	}
}

func HelpMenu() {
	fmt.Println("Here are the rules of poker and hand rankings etc")
}

func Quit() {
	fmt.Println("Thank you for playing")
}
