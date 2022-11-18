package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	sharingovercommunication "github.com/EnricoPicci/go-class/src/orchestration/count-words/sharing-over-communication"
)

func main() {
	dirPath := flag.String("dir", "", "directory containing the files")
	nummReaders := flag.Int("readers", 1, "number of concurrent readers")
	verbose := flag.Bool("verbose", false, "print the steps of the process")
	flag.Parse()

	if strings.TrimSpace(*dirPath) == "" {
		log.Fatal("The directory name is empty")
	}

	dict := sharingovercommunication.BuildDictionary(*dirPath, *nummReaders, *verbose)

	fmt.Printf("The number of files read from the directory \"%v\" is %v\n", *dirPath, len(dict.FilesRead()))
	fmt.Printf("The total number of words is %v\n", dict.TotalNumberOfWords())
	fmt.Printf("The number of unique words is %v\n", dict.NumberOfUniqueWords())
	fmt.Print("\n")
	for _, v := range dict.UniqueWords() {
		if strings.TrimSpace(v) == "" {
			fmt.Println(">>>>>>>> empty word")
		}
		fmt.Print(v + ", ")
	}
}
