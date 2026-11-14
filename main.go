package main

import (
	"fmt"
	"math/rand"
)

func main() {
	words := []string{"begin", "great", "suddenly", "before", "enough", "probably", "amaze", "exercise", "curious", "friendly", "enormous"}
	word := getWord(words)
	fmt.Println(word)
}

func getWord(words []string) string {
	if len(words) == 0 {
		return ""
	}
	index := rand.Intn(len(words)) //nolint:gosec // trade less secure pseudo random numbers for faster speeds
	word := words[index]
	return word
}
