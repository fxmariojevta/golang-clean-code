package main

import (
	"fmt"
)

func main() {
	word := "selamat malam"
	wordCount(word)
}

func wordCount(word string) {
	wordMap := make(map[string]int)

	for _, char := range word {
		stringChar := string(char)
		fmt.Println(stringChar)

		_, exist := wordMap[stringChar]
		if exist {
			wordMap[stringChar] += 1
			continue
		}
		wordMap[stringChar] = 1
	}

	fmt.Println(wordMap)
}
