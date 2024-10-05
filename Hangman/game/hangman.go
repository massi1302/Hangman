package Hangman

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	GAME_WIDTH = 155
)

func PlayHangman(word string, nbletterstoreveal int8) {
	ClearConsole()
	if nbletterstoreveal < 0 || nbletterstoreveal > 2 {
		MessageRapide("Please enter a number between 0 and 2.", 50, "rouge")
		return
	}
	lives := 10
	blanks := make([]string, len([]rune(word)))
	for i := range len([]rune(word)) {
		blanks[i] = "_"
	}
	usedLetters := make(map[rune]bool)
	var usedLettersList []string
	previousint := -1
	for i := 0; i < int(nbletterstoreveal); i++ {
		randomint := rand.Intn(len([]rune(word)) - 1)
		for previousint == randomint {
			randomint = rand.Intn(len([]rune(word)) - 1)
		}
		previousint = randomint
		blanks[randomint] = string([]rune(word)[randomint])
	}

	for lives > 0 && word != strings.Join(blanks, "") {
		time.Sleep(1 * time.Second)
		ClearConsole()

		// Display the hangman ASCII art
		hangmanArt := strings.TrimRight(Hangart[fmt.Sprintf("lifeCount%d", lives)], " \n")
		artLines := strings.Split(hangmanArt, "\n")
		maxWidth := 0
		for _, line := range artLines {
			if len(line) > maxWidth {
				maxWidth = len(line)
			}
		}

		fmt.Println()
		for _, line := range artLines {
			paddedLine := line + strings.Repeat(" ", maxWidth-len(line))
			fmt.Println(Blanc(centerString(paddedLine, GAME_WIDTH)))
		}

		// Display game status with formatting
		fmt.Println()
		heartsDisplay := strings.Repeat("❤️", lives) + strings.Repeat("💔", 10-lives)
		fmt.Println(Blanc(centerString(heartsDisplay, GAME_WIDTH)))
		//
		fmt.Println(Blanc(centerString("WORD TO GUESS :", GAME_WIDTH)))
		// Display word progress
		wordDisplay := strings.Join(blanks, " ")
		fmt.Println(Cyan(centerString(wordDisplay, GAME_WIDTH)))
		fmt.Println()

		// Display used letters
		if len(usedLettersList) > 0 {
			usedLettersStr := "Letters used: " + strings.Join(usedLettersList, ", ")
			fmt.Println(Jaune(centerString(usedLettersStr, GAME_WIDTH)))
		}

		// Input prompt - adjusted to remove extra space
		fmt.Print(Vert(centerString("Enter a letter or a word ", GAME_WIDTH/2)))

		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if len(input) > 1 {
			if word == input {
				fmt.Println()
				resultStr := fmt.Sprintf("🎉 Congratulations! You won with %d lives remaining! 🎉", lives)
				fmt.Println(Vert(centerString(resultStr, GAME_WIDTH)))
				fmt.Println(Vert(centerString(fmt.Sprintf("The word was: %s", word), GAME_WIDTH)))
				break
			} else {
				lives -= 2
				MessageRapide("❌ Incorrect guess!", 50, "rouge")
				continue
			}
		}
		if !((input >= "A" && input <= "Z") || (input >= "a" && input <= "z")) {
			MessageRapide("Please enter a single letter.", 50, "rouge")
			continue
		}

		if len(input) != 1 {
			MessageRapide("Please enter a single letter.", 50, "rouge")
			continue
		}

		letter := rune(input[0])
		if usedLetters[letter] {
			MessageRapide("You've already guessed that letter!", 50, "jaune")
			continue
		}

		usedLetters[letter] = true
		usedLettersList = append(usedLettersList, string(letter))
		correctGuess := false

		for i, wordLetter := range word {
			if letter == wordLetter {
				blanks[i] = string(letter)
				correctGuess = true
			}
		}

		if !correctGuess {
			lives--
			MessageRapide("❌ Incorrect guess!", 50, "rouge")
		} else {
			MessageRapide("✅ Correct guess!", 50, "vert")
		}
	}

	fmt.Print("Press Enter to continue...")
	fmt.Scanln()
	ClearConsole()

	// Display final result
	if word == strings.Join(blanks, "") {
		// Victory display

		fmt.Println()
		resultStr := fmt.Sprintf("🎉 Congratulations! You won with %d lives remaining! 🎉", lives)
		fmt.Println(Vert(centerString(resultStr, GAME_WIDTH)))
		fmt.Println(Vert(centerString(fmt.Sprintf("The word was: %s", word), GAME_WIDTH)))
	} else {
		// Game over display
		if word != strings.Join(blanks, "") {
			ClearConsole()
			fmt.Println()
			fmt.Println(Rouge(centerString(`
								┏━━━━━━━━━━━━━━━━━━━━━━━━━━┓
								┃      GAME  OVER !        ┃
								┗━━━━━━━━━━━━━━━━━━━━━━━━━━┛`, GAME_WIDTH)))
			fmt.Println()
			fmt.Println(Blanc(centerString("The word was:", GAME_WIDTH)))
			fmt.Println(Rouge(centerString(word, GAME_WIDTH)))
			fmt.Println()
			fmt.Println(Blanc(centerString("Press Enter to continue...", GAME_WIDTH)))
			fmt.Scanln()
		}
	}

	// Wait before returning to menu
	time.Sleep(2 * time.Second)
}
