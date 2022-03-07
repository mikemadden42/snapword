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
	index := rand.Intn(length - 1)
	word := words[index]
	fmt.Println(word)
}
