package hangman

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GetHangman(file string, indice int) {
	hangmanArray := []string{""}
	str := ""
	i := 1                  // i représente les lignes dans le fichier hangman.txt
	f, err := os.Open(file) // ouvre le fichier hangmant.txt
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f) //permet de lire par ligne le fichier txt
	for scanner.Scan() {
		line := scanner.Text() // permet d'avoir le texte de la ligne
		str += line + "\n"
		if i%8 == 0 { //le dessin du pendu est represente sur 8 lignes donc tous les 8 lignes
			hangmanArray = append(hangmanArray, str) // on mets le dessin dans une liste
			str = ""                                 //pour pouvoir l'utiliser plus facilement
		}
		i++ // on rajoute 1 à i car on change de ligne
	}
	f.Close()
	sort.Sort(sort.Reverse(sort.StringSlice(hangmanArray))) //on retourne la liste car on attemptsRemaining
	indice -= 1                                             //commence à 10 et finis à 1
	if indice > 6 {
		fmt.Println(Green + hangmanArray[indice] + Reset) //Hangman en vert
	} else if indice >= 3 {
		fmt.Println(Yellow + hangmanArray[indice] + Reset) // hangman en jaune
	} else {
		fmt.Println(Red + hangmanArray[indice] + Reset) //hangman en rouge
	}
}
func GetSave(file string) (string, string, int, []bool, string) {

	saveArray := []string{}
	boolString := ""
	found := []bool{}
	i := 0                  // i représente les lignes dans le fichier Save.txt
	f, err := os.Open(file) // ouvre le fichier Save.txt
	if err != nil {
		fmt.Println(Red + "There is no save")
		os.Exit(0)
	}
	scanner := bufio.NewScanner(f) //permet de lire par ligne le fichier txt
	for scanner.Scan() {           // permet d'avoir le texte de la ligne
		if i < 4 {
			saveArray = append(saveArray, scanner.Text()) //Récupère le mots,le stock des lettre, et le nombre d'essais restant
			i++
		} else {
			boolString = scanner.Text() // récupère found
		}

	}
	f.Close()

	foundstring := strings.Fields(boolString) //Separe string grace au espace
	for _, items := range foundstring {       //transforme les true false qui sont en string en boolean
		element, _ := strconv.ParseBool(items)
		found = append(found, element)
	}

	attempts, _ := strconv.Atoi(saveArray[3]) //transforme un string en int

	return saveArray[0], saveArray[1], attempts, found, saveArray[2]
}

func GetWords(file string) []string {
	wordsArray := []string{}
	f, err := os.Open(file) // ouvre le fichier où se trouve les mots
	if err != nil {
		fmt.Println(Red+file, "is missing")
		os.Exit(0)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		wordsArray = append(wordsArray, scanner.Text()) // mets les mots dans une liste pour pouvoir utiliser le random dessus
	}
	f.Close()

	return wordsArray
}
