package main

import (
	"testing"

	"github.com/EnricoPicci/go-class/src/testhelpers"
)

func TestMostFrequentWord(t *testing.T) {
	fileName := "testdata/count-words/file-with-repeated-words.txt"

	mostFrequent := mostFrequentWordWithCount(testhelpers.FilePath(fileName))

	// run the checks

	// check the most frequent word
	expectedMostFrequetWord := "thisIsRepeated3times"
	if expectedMostFrequetWord != mostFrequent.Word {
		t.Fatalf("Expected most frequent word \"%v\" - got \"%v\" \n", expectedMostFrequetWord, mostFrequent.Word)
	}
	// check the number of occurrences
	expectedOccurrencies := 3
	if expectedOccurrencies != mostFrequent.Occurrencies {
		t.Fatalf("Expected occurrencies \"%v\" - got \"%v\" \n", expectedOccurrencies, mostFrequent.Occurrencies)
	}
}
