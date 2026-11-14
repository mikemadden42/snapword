package main

import "testing"

func TestWord(t *testing.T) {
	words := []string{"Noah", "Isaac", "Jacob", "Abraham"}
	seen := make(map[string]bool)
	for i := 0; i < 1000; i++ {
		seen[getWord(words)] = true
	}
	for _, w := range words {
		if !seen[w] {
			t.Errorf("%q was never chosen.", w)
		}
	}
}

func TestSingleWord(t *testing.T) {
	if word := getWord([]string{"only"}); word != "only" {
		t.Errorf("Got %q, want \"only\".", word)
	}
}

func TestNoWords(t *testing.T) {
	if word := getWord(nil); word != "" {
		t.Errorf("Got %q, want an empty word.", word)
	}
}
