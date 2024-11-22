package wordle

import (
	"fmt"
	"testing"
)

func statusToString(status letterStatus) string {
	switch status {
	case none:
		return "none"
	case correct:
		return "correct"
	case present:
		return "present"
	case absent:
		return "absent"
	default:
		return "unknown"
	}
}

func TestNewWordleState(t *testing.T) {
	word := "HELLO"
	ws := newWordleState(word)
	wordleAsString := string(ws.word[:])

	for i := 0; i < len(word); i++ {
		if !(word[i] == ws.word[i]) {
			t.Logf(" lp0 on fire   ")

		}
	}

	t.Log("Created wordleState:")
	t.Logf("    word: %s", wordleAsString)
	t.Logf("    guesses: %v", ws.guesses)
	t.Logf("    currGuess: %v", ws.currGuess)
}

func TestLetter(t *testing.T) {
	var TestLetter letter
	TestLetter = newLetter(byte("H"[0]))
	if !(TestLetter.char == byte("H"[0])) {
		t.Logf(" lp0 on fire   ")

	}
	if !(TestLetter.status == 0) {
		t.Logf(" lp0 on fire   ")
	}
}

//newGuess should take in a string and return a new guess. You should loop over each letter in the string and convert them to letter structs.

func TestNewGuess(t *testing.T) {
	t.Log("Running TestNewGuess")
	testWord := "GUESS"
	TestGuess := newGuess(testWord)

	for i, v := range testWord {
		t.Logf("Comparing %c to %c. Status of guess is: %d", v, TestGuess[i].char, TestGuess[i].status)
		if byte(v) != TestGuess[i].char {
			t.Log("The chars don't match!")
		}

		if byte(v) != TestGuess[i].char {
			t.Log("The chars don't match!")
		}

		if TestGuess[i].status != 0 {
			t.Log("The guess status is not none!")
		}
	}

}

func TestAppendGuess(t *testing.T) {
	var err error

	guessWord := "YIELD"
	guess := newGuess(guessWord)

	word := "HELLO"
	ws := newWordleState(word)

	err = ws.appendGuess(guess)

	if err != nil {
		t.Fatalf("an error occurred")
	}
}

func TestUpdateLettersWithWord(t *testing.T) {
	guessWord := "YIELD"
	guess := newGuess(guessWord)

	var word [wordSize]byte
	copy(word[:], "HELLO")
	guess.updateLettersWithWord(word)

	statuses := []letterStatus{
		absent,  // "Y" is not in "HELLO"
		absent,  // "I" is not in "HELLO"
		present, // "E" is in "HELLO" but not in the correct position
		correct, // "L" is in "HELLO" and in the correct position
		absent,  // "D" is not in "HELLO"
	}

	// Check that statuses are correct
	for i, j := range statuses {
		if j != guess[i].status {
			t.Fatalf("%d is not %d", j, guess[i].status)
		}

	}
}

func TestIsWordGuessed(t *testing.T) {

	var word [wordSize]byte
	copy(word[:], "HELLO")

	guessWordIncorrect := "YIELD"
	guessIncorrect := newGuess(guessWordIncorrect)

	guessWordCorrect := "HELLO"
	guessCorrect := newGuess(guessWordCorrect)

	ws := newWordleState(guessWordCorrect)

	//Incorrect Guess
	guessIncorrect.updateLettersWithWord(word)
	ws.appendGuess(guessIncorrect)
	if ws.isWordGuessed() {
		t.Fatal("Word should be incorrect.")
	}

	//Correct Guess
	guessCorrect.updateLettersWithWord(word)
	ws.appendGuess(guessCorrect)
	if !ws.isWordGuessed() {
		t.Fatal("Word should be Correct.")
	}

}

func TestShouldEndGame(t *testing.T) {

	var word [wordSize]byte
	copy(word[:], "HELLO")

	guessWordIncorrect := "YIELD"
	guessIncorrect := newGuess(guessWordIncorrect)

	guessWordCorrect := "HELLO"
	guessCorrect := newGuess(guessWordCorrect)

	ws := newWordleState(guessWordCorrect)

	//Incorrect Guess
	guessIncorrect.updateLettersWithWord(word)
	ws.appendGuess(guessIncorrect)
	if ws.shouldEndGame() {
		t.Fatal("Game should not end")
	}

	//Correct Guess
	guessCorrect.updateLettersWithWord(word)
	ws.appendGuess(guessCorrect)
	if !ws.shouldEndGame() {
		t.Fatal("Game should end")
	}

	// Add too many
	for i := 0; i < maxGuesses; i++ {
		guessIncorrect.updateLettersWithWord(word)
		ws.appendGuess(guessIncorrect)

	}

	fmt.Printf("%d\n", ws.currGuess)
	if !ws.shouldEndGame() {
		t.Fatal("Too many guesses, game should end")
	}

}
