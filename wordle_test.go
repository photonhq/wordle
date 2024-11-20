package wordle

import "testing"

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
	var TestLetter Letter
	TestLetter = newLetter(byte("H"[0]))
	if !(TestLetter.character == byte("H"[0])) {
		t.Logf(" lp0 on fire   ")

	}
	if !(TestLetter.status == none) {
		t.Logf(" lp0 on fire   ")
	}
}
