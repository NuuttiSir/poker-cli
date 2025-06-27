package UI

import (
	"fmt"
)

func Menu() {
	fmt.Println("Welcome to my poker cli game")
	fmt.Println()
	fmt.Println("Press the corresponding number to select")
	fmt.Println("1. Play\n2. Help\n3. Quit")
}

func HelpMenu() {
	fmt.Println("Here are the rules of poker and hand rankings etc")
	return
}

func Quit() {
	fmt.Println("Thank you for playing")
}
