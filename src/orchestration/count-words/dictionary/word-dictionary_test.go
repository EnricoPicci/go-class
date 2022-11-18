package dictionary

import (
	"testing"

	"github.com/EnricoPicci/go-class/src/testhelpers"
)

func TestCountUniqueWords(t *testing.T) {
	fileName := "testdata/count-words/file-with-five-words.txt"

	fPath := testhelpers.FilePath(fileName)

	dictionary := NewWordDictionary()

	dictionary.CountUniqueWords(fPath)

	expectedNumberOfWords := 5
	gotNumberOfWords := dictionary.NumberOfUniqueWords()

	if gotNumberOfWords != expectedNumberOfWords {
		t.Errorf("The number of words in file %v is %v instead of %v", fPath, gotNumberOfWords, expectedNumberOfWords)
	}

}

func TestCountUniqueWordsWithPunctuation(t *testing.T) {
	fileName := "testdata/count-words/file-with-punctiation.txt"
	w1 := "one"
	w2 := "two"
	w3 := "three"

	fPath := testhelpers.FilePath(fileName)

	dictionary := NewWordDictionary()

	dictionary.CountUniqueWords(fPath)

	expectedNumberOfWords := 3
	gotNumberOfWords := dictionary.NumberOfUniqueWords()

	if gotNumberOfWords != expectedNumberOfWords {
		t.Errorf("The number of words in file %v is %v instead of %v", fPath, gotNumberOfWords, expectedNumberOfWords)
	}

	occurrencesOfW1 := dictionary.Occurences(w1)
	expectedOccurrencesOfW1 := 3
	if occurrencesOfW1 != expectedOccurrencesOfW1 {
		t.Errorf("The occurrences of %v in file %v are %v - expected %v", w1, fPath, occurrencesOfW1, expectedOccurrencesOfW1)
	}
	occurrencesOfW2 := dictionary.Occurences(w2)
	expectedOccurrencesOfW2 := 2
	if occurrencesOfW2 != expectedOccurrencesOfW2 {
		t.Errorf("The occurrences of %v in file %v are %v - expected %v", w2, fPath, occurrencesOfW2, expectedOccurrencesOfW2)
	}
	occurrencesOfW3 := dictionary.Occurences(w3)
	expectedOccurrencesOfW3 := 2
	if occurrencesOfW3 != expectedOccurrencesOfW3 {
		t.Errorf("The occurrences of %v in file %v are %v - expected %v", w3, fPath, occurrencesOfW3, expectedOccurrencesOfW3)
	}

	totalNumberOfWords := dictionary.TotalNumberOfWords()
	expectedTotalNumberOfWords := 7
	if totalNumberOfWords != expectedTotalNumberOfWords {
		t.Errorf("The total number of words in file %v is %v - expected %v", fPath, totalNumberOfWords, expectedTotalNumberOfWords)
	}
}

func TestCountUniqueWordsRepeatedWords(t *testing.T) {
	fileName := "testdata/count-words/file-with-repeated-words.txt"
	w1 := "thisJustOne"
	w2 := "thisIsRepeatedTwice"
	w3 := "thisIsRepeated3times"

	fPath := testhelpers.FilePath(fileName)

	dictionary := NewWordDictionary()

	dictionary.CountUniqueWords(fPath)

	expectedNumberOfWords := 5
	gotNumberOfWords := dictionary.NumberOfUniqueWords()

	if gotNumberOfWords != expectedNumberOfWords {
		t.Errorf("The number of words in file %v is %v instead of %v", fPath, gotNumberOfWords, expectedNumberOfWords)
	}

	occurrencesOfW1 := dictionary.Occurences(w1)
	expectedOccurrencesOfW1 := 1
	if occurrencesOfW1 != expectedOccurrencesOfW1 {
		t.Errorf("The occurrences of %v in file %v are %v - expected %v", w1, fPath, occurrencesOfW1, expectedOccurrencesOfW1)
	}
	occurrencesOfW2 := dictionary.Occurences(w2)
	expectedOccurrencesOfW2 := 2
	if occurrencesOfW2 != expectedOccurrencesOfW2 {
		t.Errorf("The occurrences of %v in file %v are %v - expected %v", w2, fPath, occurrencesOfW2, expectedOccurrencesOfW2)
	}
	occurrencesOfW3 := dictionary.Occurences(w3)
	expectedOccurrencesOfW3 := 3
	if occurrencesOfW3 != expectedOccurrencesOfW3 {
		t.Errorf("The occurrences of %v in file %v are %v - expected %v", w3, fPath, occurrencesOfW3, expectedOccurrencesOfW3)
	}

	totalNumberOfWords := dictionary.TotalNumberOfWords()
	expectedTotalNumberOfWords := 8
	if totalNumberOfWords != expectedTotalNumberOfWords {
		t.Errorf("The total number of words in file %v is %v - expected %v", fPath, totalNumberOfWords, expectedTotalNumberOfWords)
	}

}

func TestMerge(t *testing.T) {
	w1 := "one"
	w6 := "six"

	sourceFileName := "testdata/count-words/file-with-five-words.txt"
	targetFileName := "testdata/count-words/file-with-two-unique-words.txt"

	sourceFilePath := testhelpers.FilePath(sourceFileName)
	targetFilePath := testhelpers.FilePath(targetFileName)

	sourceDictionary := NewWordDictionary()
	targetDictionary := NewWordDictionary()

	sourceDictionary.CountUniqueWords(sourceFilePath)
	targetDictionary.CountUniqueWords(targetFilePath)

	targetDictionary.Merge(sourceDictionary)

	expectedNumberOfWords := 6
	gotNumberOfWords := targetDictionary.NumberOfUniqueWords()

	if gotNumberOfWords != expectedNumberOfWords {
		t.Errorf("After merge of file %v into %v the number of unique words is %v instead of %v",
			sourceFilePath, targetFileName, gotNumberOfWords, expectedNumberOfWords)
	}

	occurrencesOfW1 := targetDictionary.Occurences(w1)
	expectedOccurrencesOfW1 := 4
	if occurrencesOfW1 != expectedOccurrencesOfW1 {
		t.Errorf("After merge of file %v into %v the occurrences of %v iare %v - expected %v",
			sourceFilePath, targetFileName, w1, occurrencesOfW1, expectedOccurrencesOfW1)
	}

	occurrencesOfW6 := targetDictionary.Occurences(w6)
	expectedOccurrencesOfW6 := 2
	if occurrencesOfW6 != expectedOccurrencesOfW6 {
		t.Errorf("After merge of file %v into %v the occurrences of %v iare %v - expected %v",
			sourceFilePath, targetFileName, w6, occurrencesOfW6, expectedOccurrencesOfW6)
	}

	totalNumberOfWords := targetDictionary.TotalNumberOfWords()
	expectedTotalNumberOfWords := 10
	if totalNumberOfWords != expectedTotalNumberOfWords {
		t.Errorf("After merge of file %v into %v the total number of words is %v - expected %v",
			sourceFilePath, targetFileName, totalNumberOfWords, expectedTotalNumberOfWords)
	}
}

// same test as TestMerge just inverting target and source
func TestMergeInvertTargetSource(t *testing.T) {
	w1 := "one"
	w6 := "six"

	targetFileName := "testdata/count-words/file-with-five-words.txt"
	sourceFileName := "testdata/count-words/file-with-two-unique-words.txt"

	sourceFilePath := testhelpers.FilePath(sourceFileName)
	targetFilePath := testhelpers.FilePath(targetFileName)

	sourceDictionary := NewWordDictionary()
	targetDictionary := NewWordDictionary()

	sourceDictionary.CountUniqueWords(sourceFilePath)
	targetDictionary.CountUniqueWords(targetFilePath)

	targetDictionary.Merge(sourceDictionary)

	expectedNumberOfWords := 6
	gotNumberOfWords := targetDictionary.NumberOfUniqueWords()

	if gotNumberOfWords != expectedNumberOfWords {
		t.Errorf("After merge of file %v into %v the number of unique words is %v instead of %v",
			sourceFilePath, targetFileName, gotNumberOfWords, expectedNumberOfWords)
	}

	occurrencesOfW1 := targetDictionary.Occurences(w1)
	expectedOccurrencesOfW1 := 4
	if occurrencesOfW1 != expectedOccurrencesOfW1 {
		t.Errorf("After merge of file %v into %v the occurrences of %v iare %v - expected %v",
			sourceFilePath, targetFileName, w1, occurrencesOfW1, expectedOccurrencesOfW1)
	}

	occurrencesOfW6 := targetDictionary.Occurences(w6)
	expectedOccurrencesOfW6 := 2
	if occurrencesOfW6 != expectedOccurrencesOfW6 {
		t.Errorf("After merge of file %v into %v the occurrences of %v iare %v - expected %v",
			sourceFilePath, targetFileName, w6, occurrencesOfW6, expectedOccurrencesOfW6)
	}

	totalNumberOfWords := targetDictionary.TotalNumberOfWords()
	expectedTotalNumberOfWords := 10
	if totalNumberOfWords != expectedTotalNumberOfWords {
		t.Errorf("After merge of file %v into %v the total number of words is %v - expected %v",
			sourceFilePath, targetFileName, totalNumberOfWords, expectedTotalNumberOfWords)
	}
}

func TestFilesRead(t *testing.T) {
	fileName_1 := "testdata/count-words/file-with-five-words.txt"
	fileName_2 := "testdata/count-words/file-with-two-unique-words.txt"
	fileName_3 := "testdata/count-words/file-with-five-words.txt"

	fPath_1 := testhelpers.FilePath(fileName_1)
	fPath_2 := testhelpers.FilePath(fileName_2)
	fPath_3 := testhelpers.FilePath(fileName_3)

	// First file read
	dictionary := NewWordDictionary()
	dictionary.CountUniqueWords(fPath_1)

	expectedNumberOfFiles := 1
	gotNumberOfFiles := len(dictionary.FilesRead())
	if len(dictionary.FilesRead()) != expectedNumberOfFiles {
		t.Errorf("The number of files is %v instead of %v", gotNumberOfFiles, expectedNumberOfFiles)
	}
	gotExpectedFileRead_1 := dictionary.FilesRead()[0]
	if gotExpectedFileRead_1 != fPath_1 {
		t.Errorf("The file %v is found instead of %v", fPath_1, gotExpectedFileRead_1)
	}

	// Second file read
	dictionary_2 := NewWordDictionary()
	dictionary_2.CountUniqueWords(fPath_2)
	dictionary.Merge(dictionary_2)

	expectedNumberOfFiles = 2
	gotNumberOfFiles = len(dictionary.FilesRead())
	if len(dictionary.FilesRead()) != expectedNumberOfFiles {
		t.Errorf("The number of files is %v instead of %v", gotNumberOfFiles, expectedNumberOfFiles)
	}
	gotExpectedFileRead_1 = dictionary.FilesRead()[0]
	if gotExpectedFileRead_1 != fPath_1 {
		t.Errorf("The file %v is is found instead of %v", fPath_1, gotExpectedFileRead_1)
	}
	gotExpectedFileRead_2 := dictionary.FilesRead()[1]
	if gotExpectedFileRead_2 != fPath_2 {
		t.Errorf("The file %v is is found instead of %v", fPath_2, gotExpectedFileRead_2)
	}

	// Third file read
	dictionary_3 := NewWordDictionary()
	dictionary_3.CountUniqueWords(fPath_3)
	dictionary.Merge(dictionary_3)

	expectedNumberOfFiles = 3
	gotNumberOfFiles = len(dictionary.FilesRead())
	if len(dictionary.FilesRead()) != expectedNumberOfFiles {
		t.Errorf("The number of files is %v instead of %v", gotNumberOfFiles, expectedNumberOfFiles)
	}
	gotExpectedFileRead_1 = dictionary.FilesRead()[0]
	if gotExpectedFileRead_1 != fPath_1 {
		t.Errorf("The file %v is is found instead of %v", fPath_1, gotExpectedFileRead_1)
	}
	gotExpectedFileRead_2 = dictionary.FilesRead()[1]
	if gotExpectedFileRead_2 != fPath_2 {
		t.Errorf("The file %v is is found instead of %v", fPath_2, gotExpectedFileRead_2)
	}
	gotExpectedFileRead_3 := dictionary.FilesRead()[2]
	if gotExpectedFileRead_3 != fPath_1 {
		t.Errorf("The file %v is is found instead of %v", fPath_3, gotExpectedFileRead_3)
	}

}

func TestWords(t *testing.T) {
	fileName := "testdata/count-words/file-with-five-words.txt"

	fPath := testhelpers.FilePath(fileName)

	dictionary := NewWordDictionary()

	dictionary.CountUniqueWords(fPath)

	// UniqueWords returns a slice of words sorted
	expectedWords := []string{"five", "four", "one", "three", "two"}
	words := dictionary.UniqueWords()

	for i, gotWord := range words {
		expectedWord := expectedWords[i]
		if gotWord != expectedWord {
			t.Errorf("Read words from file %v. Got word %v in position %v - expected %v", fPath, gotWord, i, expectedWord)
		}
	}

}
