package Hangman

import (
	"fmt"
	"os"
	"strings"

	"github.com/eiannone/keyboard"
)

const (
	OPTION_MENU_WIDTH = 159
	OPTION_WIDTH      = 20
)

var options []string
var optionsFunc map[string]func()

// []string{"NEW GAME", "PRIVATE GAME", "SINGLE PLAYER"}
func gameInit() {
	options = []string{"EASY", "MEDIUM", "HARD", "EXIT"}
	optionsFunc = make(map[string]func())
	optionsFunc["EASY"] = func() {
		PlayHangman(RandomWord("easy_words.txt"))
	}
	optionsFunc["MEDIUM"] = func() {
		PlayHangman(RandomWord("medium_words.txt"))
	}
	optionsFunc["HARD"] = func() {
		PlayHangman(RandomWord("hard_words.txt"))
	}
	optionsFunc["EXIT"] = func() {
		ClearConsole()
		os.Exit(0)
	}
}

func AfficherMenu(wordlist string) string {
	if wordlist == "" {
		gameInit()
	} else {
		PlayHangman(RandomWord(wordlist))
	}

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		err := keyboard.Close()
		if err != nil {
			panic(err)
		}
	}()
	for {
		selectedOption := 0
		for {
			ClearConsole()
			fmt.Print("\033[2J")
			fmt.Print("\033[H")
			AfficherTitre()
			for i, option := range options {
				if i == selectedOption {
					fmt.Println(Blanc(centerString(fmt.Sprintf("┌%s┐", strings.Repeat("─", OPTION_WIDTH-2)), OPTION_MENU_WIDTH)))
					fmt.Println(Blanc(centerString(">> "+fmt.Sprintf("│%s│", centerString(option, OPTION_WIDTH-2))+" <<", OPTION_MENU_WIDTH)))
					fmt.Println(Blanc(centerString(fmt.Sprintf("└%s┘", strings.Repeat("─", OPTION_WIDTH-2)), OPTION_MENU_WIDTH)))
				} else {
					fmt.Println(Blanc(centerString(fmt.Sprintf("┌%s┐", strings.Repeat("─", OPTION_WIDTH-2)), OPTION_MENU_WIDTH)))
					fmt.Println(Blanc(centerString(fmt.Sprintf("│%s│", centerString(option, OPTION_WIDTH-2)), OPTION_MENU_WIDTH)))
					fmt.Println(Blanc(centerString(fmt.Sprintf("└%s┘", strings.Repeat("─", OPTION_WIDTH-2)), OPTION_MENU_WIDTH)))
				}
				if i < len(optionsFunc)-1 {
					fmt.Println()
				}
			}

			fmt.Println(strings.Repeat("\n", 2))

			_, key, err := keyboard.GetKey()
			if err != nil {
				panic(err)
			}

			switch key {
			case keyboard.KeyArrowUp:
				selectedOption = (selectedOption + len(optionsFunc) - 1) % len(optionsFunc)
			case keyboard.KeyArrowDown:
				selectedOption = (selectedOption + 1) % len(optionsFunc)
			case keyboard.KeyEnter:
				optionsFunc[options[selectedOption]]()
			default:
			}
		}
	}
}

func centerString(str string, width int) string {
	strArr := []rune(str)
	spaces := int(float64(width-len(strArr)) / 2)
	return strings.Repeat(" ", spaces) + str + strings.Repeat(" ", width-(spaces+len(strArr)))
}

func Blanc(str string) string {
	return "\033[37m" + str + "\033[0m"
}

func FileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}
