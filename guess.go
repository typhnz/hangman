package hangman

import (
	"bufio"
	"fmt"
	"os"
)

func Guess(word string, hangman string, file string) {
	// declaring variables
	var attempts = 10
	var found = []bool{}
	stock := ""
	// if our file is Save.txt, save the game.
	// otherwise, we pick a random letter to display and start the game
	if file == "Save.txt" {
		word, stock, attempts, found, file = GetSave(file)
	} else {
		found = make([]bool, len(word))
		for n := 0; n < len(word)/2-1; n++ {
			found[RandomLetter(word)] = true
		}
	}
	isLetter := false
	fmt.Println("\033[2J\033[1;1H")
	fmt.Println("You are now playing HANGMAN, you have 10 attempts total. Good luck!")

	// initializing a loop for each attempt
	// evlautating, printing, scanning, and storing based on the guess
	for attemptsRemaining := attempts; attemptsRemaining <= 11 && attemptsRemaining >= 1; attemptsRemaining-- {
		Prints(word, found) // displays platorm with letters or _
		fmt.Printf(" You have %v attempts. \n", attemptsRemaining)
		fmt.Print("Guess: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		guess := scanner.Text()           // taking user input
		guess, isLetter = Isletter(guess) // evaluating if user input is a letter
		if !isLetter {
			fmt.Println("Enter a valid letter: ")
		} else {
			Save(guess, word, stock, attemptsRemaining, found, file)         // save game
			found, validGuess, validAll, isWord := Found(word, guess, found) // reinitializing boolean variables based on Found()
			fmt.Println("\033[2J\033[1;1H")
			if attemptsRemaining == 1 && !Success(found) { // case: loss
				fmt.Println("So sad! You lost!")
				fmt.Printf("The answer is: %v\n", word)
				break
			} else if Success(found) && validAll { // case: win
				attemptsRemaining++
				fmt.Println("Amazing! You won!")
				Prints(word, found)
				AddWord(file)
				break
			}
			if validGuess { // case: valid single letter guess
				attemptsRemaining++
				fmt.Println(" You got it! Keep going!")
			} else if !isWord { // case: invalid single letter guess & not a word
				fmt.Println(" Too bad! Try again!")
			} else {
				attemptsRemaining-- // case: invalid single letter guess
				fmt.Println(" Too bad! Try again!")
			}
			stock = stockGuess(guess, stock) // store guess in array
			fmt.Println(stock)
		}
		GetHangman(hangman, attemptsRemaining) // display hangman based on attempts
		fmt.Println(stock)
	}
}
