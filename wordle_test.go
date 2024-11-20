package wordle

import "testing"

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
