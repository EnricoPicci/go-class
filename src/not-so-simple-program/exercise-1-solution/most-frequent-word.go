package main

import "github.com/EnricoPicci/go-class/src/not-so-simple-program/helpers"

// mostFrequentWordWithCount returns a value of type helpers.WordOccurrency that contains the word which is most frequent in the file read as well as
// the number of occurencies of that word
func mostFrequentWordWithCount(file string) helpers.WordOccurrency {
	var mostFrequent helpers.WordOccurrency

	wordOccurrencies := helpers.FileWordOccurrencies(file)
	var maxOccurrencies int
	for i := 0; i < len(wordOccurrencies); i++ {
		occurrencies := wordOccurrencies[i].Occurrencies
		if occurrencies > maxOccurrencies {
			maxOccurrencies = occurrencies
			mostFrequent = wordOccurrencies[i]
		}
	}

	return mostFrequent
}
