package helpers

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/EnricoPicci/go-class/src/testhelpers"
)

func TestWordOccurrencies(t *testing.T) {
	// data is a string with 7 words, out of which 5 are different
	data := []byte("these are five five different words words")
	// convert byte slice to io.Reader
	reader := bytes.NewReader(data)
	// create the scanner
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	wordOccurrencies := wordOccurrencies(scanner)

	// run the checks

	// check that there are 5 unique words
	expectedNumOfWords := 5
	gotNumeOfWOrds := len(wordOccurrencies)
	if expectedNumOfWords != gotNumeOfWOrds {
		t.Fatalf("Expected num of words %v - got %v\n", expectedNumOfWords, gotNumeOfWOrds)
	}

	var checkedWord string
	var expectedOccurrencies int
	var gotOccurrencies int
	// check that the word five occurs 2 times
	checkedWord = "five"
	for i := range wordOccurrencies {
		if wordOccurrencies[i].Word == checkedWord {
			gotOccurrencies = wordOccurrencies[i].Occurrencies
		}
	}
	expectedOccurrencies = 2
	if expectedOccurrencies != gotOccurrencies {
		t.Fatalf("Expected occurrencies of word \"%v\" %v - got %v\n", checkedWord, expectedOccurrencies, gotOccurrencies)
	}

	// check that the word "different" occurs 1 time
	checkedWord = "different"
	for i := range wordOccurrencies {
		if wordOccurrencies[i].Word == checkedWord {
			gotOccurrencies = wordOccurrencies[i].Occurrencies
		}
	}
	expectedOccurrencies = 1
	if expectedOccurrencies != gotOccurrencies {
		t.Fatalf("Expected occurrencies of word \"%v\" %v - got %v\n", checkedWord, expectedOccurrencies, gotOccurrencies)
	}
}

func TestFileWordOccurrencies(t *testing.T) {
	fileName := "testdata/count-words/file-with-five-words.txt"

	wordOccurrencies := FileWordOccurrencies(testhelpers.FilePath(fileName))

	// run the checks

	// check that there are 5 unique words
	expectedNumOfWords := 5
	gotNumOfWOrds := len(wordOccurrencies)
	if expectedNumOfWords != gotNumOfWOrds {
		t.Fatalf("Expected num of words %v - got %v\n", expectedNumOfWords, gotNumOfWOrds)
	}
}
