package hangman

import (
	"math/rand"
	"time"
)

func Random(words []string) string {
	rand.Seed(time.Now().UnixNano())      //le nombre de seconde depuis le 1er janvier 1970
	return words[rand.Intn(len(words)-1)] //retourne un mots aléatoire de la liste
}

func RandomLetter(words string) int {
	rand.Seed(time.Now().UnixNano()) //le nombre de seconde depuis le 1er janvier 1970
	return rand.Intn(len(words))     // retourne une lettre aléatoire à afficher au début de la partie
}
