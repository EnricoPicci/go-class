package main

import "github.com/EnricoPicci/go-class/src/not-so-simple-program/helpers"

func mostFrequentWord(file string) string {
	wordOccurrencies := helpers.FileWordOccurrencies(file)

	var mostFrequent string
	var maxOccurrencies int

	for i := 0; i < len(wordOccurrencies); i++ {
		occurrencies := wordOccurrencies[i].Occurrencies
		if occurrencies > maxOccurrencies {
			maxOccurrencies = occurrencies
			mostFrequent = wordOccurrencies[i].Word
		}
	}

	return mostFrequent
}
