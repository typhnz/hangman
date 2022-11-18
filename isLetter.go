package hangman

func Isletter(guess string) (string, bool) { //permet de changer toutes les lettres qui sont des caractere
	guessRune := []rune(guess) // spéciaux en lettre miniscule

	for i := 0; i < len(guessRune); i++ {
		if guessRune[i] <= 'z' && guessRune[i] >= 'a' {
		} else if guessRune[i] <= 'Z' && guessRune[i] >= 'A' {
			guessRune[i] += 32
		} else if guessRune[i] <= 'å' && guessRune[i] >= 'à' {
			guessRune[i] = ('a')
		} else if guessRune[i] == 'ç' {
			guessRune[i] = ('c')
		} else if guessRune[i] <= 'ë' && guessRune[i] >= 'è' {
			guessRune[i] = ('e')
		} else if guessRune[i] <= 'ï' && guessRune[i] >= 'ì' {
			guessRune[i] = ('i')
		} else if guessRune[i] <= 'ö' && guessRune[i] >= 'ò' {
			guessRune[i] = ('o')
		} else if guessRune[i] == 'ñ' {
			guessRune[i] = ('n')
		} else if guessRune[i] <= 'ü' && guessRune[i] >= 'ù' {
			guessRune[i] = ('u')
		} else {
			return guess, false //renvoie une erreur si un caractere est autre chose que une lettre
		}
	}

	return string(guessRune), true
}

// 224-229 A
// 231 c
//232-235 e
//236-239 i
//241 n
//242-246 o
