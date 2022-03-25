package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	words := []string{"begin", "great", "suddenly", "before", "enough", "probably"}
	rand.Seed((time.Now().UnixNano()))
	length := len(words)
	index := rand.Intn(length - 1) //nolint:gosec // trade less secure pseudo random numbers for faster speeds
	word := words[index]
	fmt.Println(word)
}
