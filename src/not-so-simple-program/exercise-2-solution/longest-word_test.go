package main

import (
	"testing"

	"github.com/EnricoPicci/go-class/src/testhelpers"
)

func TestLongestWord(t *testing.T) {
	fileName := "testdata/count-words/file-with-five-words.txt"

	longest := longestWord(testhelpers.FilePath(fileName))

	// run the checks

	// check the most frequent word
	expectedLongest := "three"
	if expectedLongest != longest {
		t.Fatalf("Expectedlongest word \"%v\" - got \"%v\" \n", expectedLongest, longest)
	}
}
