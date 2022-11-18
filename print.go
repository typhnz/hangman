package hangman

import (
	"fmt"
)

func Prints(word string, found []bool) {
	runeWord := []rune(word)             // converting our word -string to runes
	for i := 0; i < len(runeWord); i++ { // looping through the word in rune form
		if found[i] { // case: letter is found/true at any given index
			fmt.Print(string(runeWord[i])) // letter is printed/shown at corresponding index
		} else {
			fmt.Print("_ ") // letter is not found, so print _ at index
		}
	}
	fmt.Print("\n")
}
