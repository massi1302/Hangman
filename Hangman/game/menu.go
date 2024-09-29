package Hangman

import (
	"fmt"
	"os"
	"strings"

	"github.com/eiannone/keyboard"
)

const (
	UP    = keyboard.KeyArrowUp
	DOWN  = keyboard.KeyArrowDown
	ENTER = keyboard.KeyEnter
)

func AfficherMenu() int {
	for {

		// Options de menu
		options := []string{"   NEW GAME  ", " PRIVATE GAME", "SINGLE PLAYER"}
		selectedOption := 0
		if err := keyboard.Open(); err != nil {
			panic(err)
		}
		defer keyboard.Close()
		for {
			fmt.Print("\033[2J")
			fmt.Print("\033[H")

			AfficherTitre()

			for i, option := range options {
				if i == selectedOption {
					fmt.Println(Blanc(Center(fmt.Sprintf("                                                                       ┌─────────────────┐"), 80)))
					fmt.Println(Blanc(Center(fmt.Sprintf("                                                                    >> │ %s   │ <<", option), 80)))
					fmt.Println(Blanc(Center(fmt.Sprintf("                                                                       └─────────────────┘"), 80)))
				} else {
					fmt.Println(Blanc(Center("                                                                       ┌─────────────────┐", 80)))
					fmt.Println(Blanc(Center(fmt.Sprintf("                                                                       │ %s   │", option), 80)))
					fmt.Println(Blanc(Center("                                                                       └─────────────────┘", 80)))
				}
				if i < len(options)-1 {
					fmt.Println()
				}
			}

			fmt.Println(strings.Repeat("\n", 2))

			char, key, err := keyboard.GetKey()
			if err != nil {
				panic(err)
			}

			switch key {
			case UP:
				if selectedOption > 0 {
					selectedOption--
				}
			case DOWN:
				if selectedOption < len(options)-1 {
					selectedOption++
				}
			case ENTER:
				return selectedOption + 1
			case keyboard.KeyEsc:
				os.Exit(0)
			default:
				if char == 'q' || char == 'Q' {
					os.Exit(0)
				}
			}
		}
	}
}

// Center centers a string within a given width.
func Center(s string, width int) string {
	if len(s) >= width {
		return s
	}
	padding := (width - len(s)) / 2
	return strings.Repeat(" ", padding) + s + strings.Repeat(" ", width-len(s)-padding)
}

func Titre() {
	AfficherMenu()
}

// Fonction pour créer du texte blanc
func Blanc(texte string) string {
	return "\033[37m" + texte + "\033[0m"
}
