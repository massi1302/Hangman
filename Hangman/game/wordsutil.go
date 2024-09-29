package Hangman

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

// ficher ou l'on lit les ficher txt
func RandomWord(filename string) string {
	rand.Seed(time.Now().UnixNano())
	L := Reader(filename)
	return L[rand.Intn(len(L))]
}

func Reader(filename string) []string {
	res := []string{}
	content, error := ioutil.ReadFile("Data\\" + filename)
	if error != nil {
		fmt.Println("Error when opening file", error)

	}
	temp := ""
	for _, element := range string(content) {
		if element != '\n' && element != '\r' {
			temp += string(element)
		} else {
			res = append(res, temp)
			temp = ""
		}
	}
	return res
}
