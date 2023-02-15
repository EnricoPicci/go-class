package main

import "github.com/EnricoPicci/go-class/src/not-so-simple-program/helpers"

// longestWord returns the longest word found in the file.
func longestWord(file string) string {
	var longestWord string

	wordOccurrencies := helpers.FileWordOccurrencies(file)
	var maxLength int
	for i := 0; i < len(wordOccurrencies); i++ {
		word := wordOccurrencies[i].Word
		wordLength := len(word)
		if wordLength > maxLength {
			maxLength = wordLength
			longestWord = word
		}
	}

	return longestWord
}
