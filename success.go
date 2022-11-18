package hangman

func Success(found []bool) bool {
	for i := range found { // loop through boolean values
		if found[i] == false { // case: if bool is false at any given index
			return false // success !found
		}
	}
	return true // success found
}
