package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strings"

	communicationoversharing "github.com/EnricoPicci/go-class/src/orchestration/count-words/communication-over-sharing"
)

func main() {
	dirPath := flag.String("dir", "", "directory containing the files")
	numReaders := flag.Int("readers", 1, "number of concurrent readers")
	printWords := flag.Bool("printWords", false, "print the unique words found")
	numWords := flag.Int("numWords", math.MaxInt, "number of unique words to print")
	byOccurrencies := flag.Bool("byOccurrencies", false, "print the unique words found sorted by occurrencies")
	verbose := flag.Bool("verbose", false, "print the steps of the process")
	flag.Parse()

	if strings.TrimSpace(*dirPath) == "" {
		log.Fatal("The directory name is empty")
	}

	dict := communicationoversharing.BuildDictionary(*dirPath, *numReaders, *verbose)

	fmt.Printf("Files read from the directory %v\n", *dirPath)
	dict.PrintData(*printWords, *byOccurrencies, *numWords)
}
