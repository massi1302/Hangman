package Hangman

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
)

func ReadWordList(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func ChooseRandomWord(words []string) string {
	return words[rand.Intn(len(words))]
}
