package main

import (
	hangman "Hangman/game"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fichier := ""
	nbletterstoreveal := int64(0)
	args := os.Args[1:]
	if len(args) == 1 {
		fichier = args[0]
	} else if len(args) == 2 {
		fichier = args[0]

		nbletterstoreveal, _ = strconv.ParseInt(args[1], 10, 8)
	}
	if fichier != "" && !hangman.FileExists(fichier) {
		fmt.Println("Fatal: The file " + fichier + " does not exist.")
		return
	}
	hangman.AfficherMenu(fichier, int8(nbletterstoreveal))
}
