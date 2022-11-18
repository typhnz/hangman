package main

import (
	"hangman"
)

func main() {
	file := hangman.DisplayHome()
	if file == "Save.txt" {
		hangman.Guess("Save", "hangman.txt", "Save.txt")

	} else {
		wordsArray := hangman.GetWords(file)
		words := hangman.Random(wordsArray)
		hangman.Guess(words, "hangman.txt", file)
	}
}
