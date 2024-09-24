package main

import (
	Hangman "Hangman/game"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Hangman!")
	rand.Seed(time.Now().UnixNano())
	Hangman.PlayHangman()
}
