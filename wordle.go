package wordle

import (
	"fmt"
	"wordle/words"
)

const (
	maxGuesses = 6
	wordSize   = 5
)

type wordleState struct {
	// word is the word that the user is trying to guess
	word [wordSize]byte
	// guesses holds the guesses that the user has made
	guesses [maxGuesses]guess
	// currGuess is the index of the available slot in guesses
	currGuess int
}

// guess is an attempt to guess the word
type guess [wordSize]letter

type letter struct {
	// char is the letter that this struct represents
	char byte
	// status is the state of the letter (absent, present, correct)
	status letterStatus
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

// newLetter builds a new letter from a byte
func newLetter(char byte) letter {
	return letter{char: char, status: none}
}

// newGuess builds a new guess from a string
func newGuess(guessedWord string) guess {
	guess := guess{}
	for i, c := range guessedWord {
		guess[i] = newLetter(byte(c))
	}

	return guess
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
