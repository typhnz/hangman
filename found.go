package hangman

func Found(word string, guess string, found []bool) ([]bool, bool, bool, bool) {
	guessRune := []rune(guess)
	wordRune := []rune(word)
	validGuess := false
	validAll := true
	isWord := false

	if len(guess) <= 1 { // case: guess is a single letter
		for i, element := range word { // loop through the word and compare the letter
			if guessRune[0] == element { // case: letter has been found at any given index in the word
				found[i] = true   //case: boolean value is true
				validGuess = true // case: guess is valid
			}
		}
	} else if len(guess) > 1 { // case: guess is also a word
		isWord = true
		if len(guess) == len(word) {
			for j := 0; j < len(guess); j++ { // loop through the word and compare the guess
				if wordRune[j] != guessRune[j] { // case: letter in word != letter in guess
					validAll = false // case: whole giess is invalid
				}
			}
			if validAll == true { // case: whole guess is valid
				for k := 0; k < len(found); k++ { // loop through the word
					found[k] = true // case: boolean value is true
				}
			}
		} else {
			validAll = false
		}
	}

	return found, validGuess, validAll, isWord
}
