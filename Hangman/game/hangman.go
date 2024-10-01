package Hangman

import (
	"fmt"
	"strings"
)

func PlayHangman(word string) {
	ClearConsole()
	lives := 10
	blanks := make([]string, len([]rune(word)))
	for i := range len([]rune(word)) {
		blanks[i] = "_"
	}
	usedLetters := make(map[rune]bool)
	for lives > 0 && word != strings.Join(blanks, "") {
		ClearConsole()
		fmt.Printf("\n❤️ %d, Word: %s\n", lives, strings.Join(blanks, " "))
		fmt.Print("Enter a letter: ")
		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if len(input) != 1 {
			fmt.Println("Please enter a single letter.")
			continue
		}
		letter := rune(input[0])
		if usedLetters[letter] {
			fmt.Println("You've already guessed that letter.")
			continue
		}
		usedLetters[letter] = true
		correctGuess := false
		for i, wordLetter := range word {
			if letter == wordLetter {
				blanks[i] = string(letter)
				correctGuess = true
			}
		}
		if !correctGuess {
			lives--
			fmt.Print("Incorrect guess!")
		} else {
			fmt.Print("Correct guess!")
		}
	}
	if word == strings.Join(blanks, "") {
		fmt.Println(word)
		fmt.Println(strings.Join(blanks, ""))
		fmt.Printf("\n❤️ %d, Word: %s - You won, congrats!\n", lives, word)
	} else {
		fmt.Printf("\n❤️ 0, Word: %s - Sorry, you lost! \n", Hangart["lifeCount0"])
	}
}
