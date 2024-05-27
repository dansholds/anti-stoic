package antistoic

import (
	"testing"
)

func TestGetRandomQuote(t *testing.T) {
	quote := GetRandomQuote()
	if len(quote) == 0 {
		t.Error("Expected a non-empty quote")
	}
}
