package helpers

import (
	"bufio"
	"log"
	"os"
)

type WordOccurrency struct {
	Word         string
	Occurrencies int
}

// https://gosamples.dev/read-file/#read-a-file-word-by-word
func FileWordOccurrencies(filePath string) []WordOccurrency {
	// open file
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file word by word using scanner
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	return wordOccurrencies(scanner)
}

func wordOccurrencies(scanner *bufio.Scanner) []WordOccurrency {
	wordDict := map[string]int{}
	for scanner.Scan() {
		word := scanner.Text()
		occurrencies, found := wordDict[word]
		if !found {
			wordDict[word] = 0
		}
		wordDict[word] = occurrencies + 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wordOccurencies := make([]WordOccurrency, len(wordDict))
	i := 0
	for k, v := range wordDict {
		wordOccurencies[i] = WordOccurrency{k, v}
		i++
	}

	return wordOccurencies
}
