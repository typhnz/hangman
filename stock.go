package hangman

import "fmt"

func stockGuess(guess string, stock string) string {
	inArray := false     // initializing bool to tell if guess is in the stock
	if len(guess) <= 1 { // case: guess is a single letter
		if len(stock) == 0 { // case: stock is empty
			stock += guess // case: guess added to stock
		} else { // case: stock is not empty
			for _, letter := range stock { // loop through the stock
				if string(letter) == guess { // case: guess in found in stock
					fmt.Printf("ERROR!!! You have already guessed %v\n", guess)
					inArray = true // case: guess is already in stock
				}
			}
			if !inArray { // case: guess is not in stock
				stock += guess // case: guess is added to stock
			}
		}
	}
	return stock
}
