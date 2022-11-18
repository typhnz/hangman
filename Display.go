package hangman

import (
	"bufio"
	"fmt"
	"os"
)

func DisplayHelp() {
	fmt.Println("\033[2J\033[1;1H")
	fmt.Println(Cyan, "<--------------------------------------------------------Help-------------------------------------------------------->", Reset)
	fmt.Println("You can enter any letter of the alphabet , accented letters are accepted, if the letter is incorrect, your attempts decrease by 1 ")
	fmt.Println("You can enter an entire word, if the entire word is incorrect, your attempts decrease by 2")
	fmt.Println("You can enter ", Red, "stop", Reset, " at any moment while you guess the word")
	fmt.Println("That will save the game, and you can relaunch the saved game ")
	fmt.Println("Enter", Green, "Home", Reset, "to go back in the Home Menu")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	string := scanner.Text()
	if string != "Home" {
		DisplayHelp()
	}
}

func DisplayHome() string {
	fmt.Println("\033[2J\033[1;1H")
	fmt.Println(Cyan + "Welcome to Hangman" + Reset)
	fmt.Println(Green + "1  Easy" + Reset)
	fmt.Println(Yellow + "2  Medium" + Reset)
	fmt.Println(Red + "3  Hard" + Reset)
	fmt.Println(Purple + "4  Save" + Reset)
	fmt.Println(Cyan + "5  Help" + Reset)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	file := scanner.Text()
	switch file {
	case "1": //niveaux facile
		return "words.txt"
	case "2": //niveaux moyen
		return "words2.txt"
	case "3": //niveaux difficile
		return "words3.txt"
	case "4": //Sauvegarde
		return "Save.txt"
	case "5":
		DisplayHelp()
	default: //si il entre autre chose que 1,2,3 ou 4 il rappelle la fonction
		fmt.Println("Enter a correct Number")
		return DisplayHome()
	}
	return DisplayHome()
}
