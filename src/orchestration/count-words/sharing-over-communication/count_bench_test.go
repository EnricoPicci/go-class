package sharingovercommunication

import (
	"testing"

	"github.com/EnricoPicci/go-class/src/testhelpers"
)

var result int

func BenchmarkBuildDictionary_1_Reader(b *testing.B) {
	dirName := "canti-divina-commedia"
	numberOfReaders := 1

	dirPath := testhelpers.FilePath(dirName)

	var gotNumberOfWords int
	bN := b.N
	for i := 0; i < bN; i++ {
		gotNumberOfWords += CountUniqueWords(dirPath, numberOfReaders, false)
	}
	result = gotNumberOfWords
}

func BenchmarkBuildDictionary_10_Readers(b *testing.B) {
	dirName := "canti-divina-commedia"
	numberOfReaders := 10

	dirPath := testhelpers.FilePath(dirName)

	var gotNumberOfWords int
	bN := b.N
	for i := 0; i < bN; i++ {
		gotNumberOfWords += CountUniqueWords(dirPath, numberOfReaders, false)
	}
	result = gotNumberOfWords
}

func BenchmarkBuildDictionary_100_Readers(b *testing.B) {
	dirName := "canti-divina-commedia"
	numberOfReaders := 100

	dirPath := testhelpers.FilePath(dirName)

	var gotNumberOfWords int
	bN := b.N
	for i := 0; i < bN; i++ {
		gotNumberOfWords += CountUniqueWords(dirPath, numberOfReaders, false)
	}
	result = gotNumberOfWords
}
