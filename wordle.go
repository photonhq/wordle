package wordle

import "fmt"

type wordleState struct {
	// word is the word that the user is trying to guess
	word [5]byte
	// guesses holds the guesses that the user has made
	guesses [6]string
	// currGuess is the index of the available slot in guesses
	currGuess int
}

func wordle() {
	fmt.Println("welcome to wordle!")
	var myWorld wordleState
	myWorld.currGuess = 1
	fmt.Println(myWorld.currGuess)
}
