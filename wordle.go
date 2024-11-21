package wordle

import (
	"errors"
	"fmt"
	"log"
	"strings"
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

// updateLettersWithWord updates the status of the letters in the guess based on a word
func (g *guess) updateLettersWithWord(word [wordSize]byte) {

	for i, _ := range g {
		// c.char is a letter

		if g[i].char == word[i] && g[i].status == none {
			g[i].status = correct
		}

		for _, v := range word {
			if v == g[i].char && g[i].status == none {
				g[i].status = present

			}
		}
		if g[i].status == none {
			g[i].status = absent
		}
	}

}

func guessToString(g guess) string {
	var s strings.Builder
	for i := 0; i < len(g); i++ {
		s.WriteString(string(g[i].char))
	}
	return s.String()
}

// appendGuess adds a guess to the wordleState. It returns an error
// if the guess is invalid.
func (w *wordleState) appendGuess(g guess) error {
	var err error

	if w.currGuess < maxGuesses {
		if len(g) != wordSize {
			fmt.Printf("Your guess is not of lenth %d\n", wordSize)
			err = errors.New("Guess is not correct length")
		}

		if !words.IsWord(guessToString(g)) {
			fmt.Println("Your guess is not a valid word")
			err = errors.New("Guess is not a valid word")
		}
		w.currGuess++

	} else {
		log.Fatalf("Game over :(")
		err = errors.New("Too many guesses")
	}

	return err

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
