package main

import "testing"

func TestWord(t *testing.T) {
	words := []string{"Noah", "Isaac", "Jacob", "Abraham"}
	word := getWord(words)
	if word == "" {
		t.Error("Got an empty word.")
	}
}
