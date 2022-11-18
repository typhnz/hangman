package hangman

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

//Permet de rajouter un mots a la base de donné de la dificulté
func AddWord(file string) {
	word := ""
	fmt.Println("Enter a Word for the DataBase")
	fmt.Scan(&word)
	word, valid := Isletter(word) //Permet de vérifier si le mot a des caractere spéciaux
	for !valid {                  //si il possede des caracteres speciaux
		fmt.Println("Enter a Word with only letters:") //qui sont différents de lettre
		fmt.Scan(&word)                                //alors on lui redemande de réecrire le mot
		word, valid = Isletter(word)
	}
	word = "\n" + word

	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0600) //Ouvre le fichier de la difficulté choisi
	if err != nil {
		panic(err)
	}

	if _, err = f.WriteString(word); err != nil { //ecrit le mot a la suite dans le fichier
		panic(err)
	}
	f.Close()
	os.Exit(0) // arrete le programme
}

func Save(guess string, word string, stock string, attemptsRemaining int, found []bool, file string) {
	if guess == "stop" {

		//Creation du fichier et ajout du mot
		wordByte := []byte(word + "\n")
		err := ioutil.WriteFile("Save.txt", wordByte, 0777)
		if err != nil {
			fmt.Print(err)
		}
		//ouverture du fichier pour pouvoir ecrire dedans
		f, err := os.OpenFile("Save.txt", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		//Ecrie dans le fichier
		//Stock
		if _, err = f.WriteString(stock + "\n"); err != nil {
			panic(err)
		}

		//File
		if _, err = f.WriteString(file + "\n"); err != nil {
			panic(err)
		}

		//attemptsRemaining
		attemptsRemainingString := strconv.Itoa(attemptsRemaining)
		if _, err = f.WriteString(attemptsRemainingString + "\n"); err != nil {
			panic(err)
		}

		//found
		for _, element := range found {
			boolString := strconv.FormatBool(element)
			if _, err = f.WriteString(boolString + " "); err != nil {
				panic(err)
			}
		}

		f.Close()

		fmt.Println(Green + "Game Saved")
		os.Exit(0)
	}
}
