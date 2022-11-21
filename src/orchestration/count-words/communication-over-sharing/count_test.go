package communicationoversharing

import (
	"testing"

	"github.com/EnricoPicci/go-class/src/testhelpers"
)

func TestCountUniqueWords(t *testing.T) {
	dirName := "testdata/count-words"
	numberOfReaders := 10

	dirPath := testhelpers.FilePath(dirName)

	expectedNumberOfWords := 11
	gotNumberOfWords := CountUniqueWords(dirPath, numberOfReaders, true)

	if gotNumberOfWords != expectedNumberOfWords {
		t.Errorf("The number of words in all files of dir %v is %v instead of %v", dirPath, gotNumberOfWords, expectedNumberOfWords)
	}

}

func TestCountTotalNumberOfWords(t *testing.T) {
	dirName := "testdata/count-words"
	numberOfReaders := 10

	dirPath := testhelpers.FilePath(dirName)

	expectedNumberOfWords := 25
	gotNumberOfWords := CountTotalNumberOfWords(dirPath, numberOfReaders, false)

	if gotNumberOfWords != expectedNumberOfWords {
		t.Errorf("The total number of words in all files of dir %v is %v instead of %v", dirPath, gotNumberOfWords, expectedNumberOfWords)
	}

}
