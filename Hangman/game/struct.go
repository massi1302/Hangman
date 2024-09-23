package Hangman

import (
	"fmt"
	"math/rand"
	"strings"
)

type Game struct {
	Word            string
	GuessedLetters  map[rune]bool
	RemainingTries  int
	RevealedLetters []rune
}

func NewGame(word string, initialTries int) *Game {
	revealed := make([]rune, len(word))
	for i := range revealed {
		revealed[i] = '_'
	}

	// Reveal random letters (e.g., 1/4 of the word)
	numToReveal := len(word) / 4
	for i := 0; i < numToReveal; {
		idx := rand.Intn(len(word))
		if revealed[idx] == '_' {
			revealed[idx] = rune(word[idx])
			i++
		}
	}

	return &Game{
		Word:            word,
		GuessedLetters:  make(map[rune]bool),
		RemainingTries:  initialTries,
		RevealedLetters: revealed,
	}
}

func (g *Game) GuessLetter(letter rune) bool {
	letter = rune(strings.ToLower(string(letter))[0])

	if g.GuessedLetters[letter] {
		return false
	}

	g.GuessedLetters[letter] = true

	if strings.ContainsRune(g.Word, letter) {
		for i, char := range g.Word {
			if char == letter {
				g.RevealedLetters[i] = letter
			}
		}
		return true
	}

	g.RemainingTries--
	return false
}

func (g *Game) GuessWord(word string) bool {
	if word == g.Word {
		for i, char := range g.Word {
			g.RevealedLetters[i] = char
		}
		return true
	}
	g.RemainingTries -= 2
	return false
}

func (g *Game) IsWon() bool {
	return string(g.RevealedLetters) == g.Word
}

func (g *Game) IsLost() bool {
	return g.RemainingTries <= 0
}

func (g *Game) Display() {
	fmt.Printf("Word: %s\n", string(g.RevealedLetters))
	fmt.Printf("Tries left: %d\n", g.RemainingTries)
	fmt.Printf("Guessed letters: %v\n", g.GuessedLetters)
}
