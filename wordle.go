package wordle

import (
	"fmt"

	"github.com/photonhq/wordle/words"
)

type wordleState struct {
	// word is the word that the user is trying to guess
	word [5]byte
	// guesses holds the guesses that the user has made
	guesses [6]string
	// currGuess is the index of the available slot in guesses
	currGuess int
}

// letterStatus can be none, correct, present, or absent
type letterStatus int

const (
	// none = no status, not guessed yet
	none letterStatus = iota
	// absent = not in the word
	absent
	// present = in the word, but not in the correct position
	present
	// correct = in the correct position
	correct
)

// newWordleState builds a new wordleState from a string.
// Pass in the word you want the user to guess.
func newWordleState(word string) wordleState {
	var myWordle wordleState

	for i := 0; i < len(word); i++ {
		myWordle.word[i] = byte(word[i])
	}
	return myWordle
}

func wordle() {
	fmt.Println("welcome to wordle!")
	newWord := words.GetWord()
	myWordleState := newWordleState(newWord)
	fmt.Println(myWordleState.word)
}

func main() {
	wordle()
}
