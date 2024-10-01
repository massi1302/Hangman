package Hangman

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func RandomWord(filename string) string {
	words := FileToStringArray(filename)
	randomIndex := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(words))
	return words[randomIndex]
}

func FileToStringArray(filename string) []string {
	var res []string
	content, err := os.ReadFile("Data\\" + filename)
	if err != nil {
		fmt.Println("Error when opening file", err)
	}
	temp := ""
	for _, element := range string(content) {
		if (element >= 'A' && element <= 'Z') || (element >= 'a' && element <= 'z') {
			temp += string(element)
		} else {
			if len(temp) != 0 {
				res = append(res, temp)
				temp = ""
			}
		}
	}
	return append(res, temp)
}
