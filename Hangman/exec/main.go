package main

import (
	Hangman "Hangman/game"
	"math/rand"

	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	Hangman.Titre()
	Hangman.PlayHangman(Hangman.Reader("words.txt"))
}
