package main

import (
	"testing"

	"github.com/EnricoPicci/go-class/src/testhelpers"
)

func TestMostFrequentWord(t *testing.T) {
	fileName := "testdata/count-words/file-with-repeated-words.txt"

	mostFrequent := mostFrequentWord(testhelpers.FilePath(fileName))

	// run the checks

	// check the most frequent word
	expectedMostFrequetWord := "thisIsRepeated3times"
	if expectedMostFrequetWord != mostFrequent {
		t.Fatalf("Expected most frequent word \"%v\" - got \"%v\" \n", expectedMostFrequetWord, mostFrequent)
	}
}
