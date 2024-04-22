package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("Welcome to guess Game")

	// rand.Seed(time.Now().UnixNano())
	randIdx := rand.Intn(len(words()))

	wordToGuess := (words()[randIdx])
	
	fmt.Println(wordToGuess)
}
