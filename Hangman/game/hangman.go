package Hangman

import (
	"fmt"
	"strings"
)

func PlayHangman(words []string) {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")

	lives := 10

	blanks := make([]string, len(words))
	for i := range blanks {
		blanks[i] = "_"
	}

	usedLetters := make(map[rune]bool)

	for {

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
		for i, wordLetter := range words {
			if letter == rune(wordLetter[0]) {
				blanks[i] = string(letter)
				correctGuess = true
			}
		}

		if !correctGuess {
			lives--
			fmt.Println("Incorrect guess!")
		} else {
			fmt.Println("Correct guess!")
		}

		if lives <= 0 {

			fmt.Printf("\n❤️ 0, Word: %s - Sorry, you lost! \n", Hangart["lifeCount0"])
			break
		}

		if strings.Join(words, "") == strings.Join(blanks, "") {
			fmt.Printf("\n❤️ %d, Word: %s - You won, congrats!\n", lives, words)
			break
		}
	}
}
