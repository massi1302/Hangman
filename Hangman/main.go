package main

import (
	hangman "Hangman/game"
	"fmt"
	"os"
)

func main() {
	fichier := ""
	args := os.Args[1:]
	if len(args) == 1 {
		fichier = args[0]
	}

	if fichier != "" && !hangman.FileExists(fichier) {
		fmt.Println("Fatal: The file " + fichier + " does not exist.")
		return
	}
	hangman.AfficherMenu(fichier)
}
