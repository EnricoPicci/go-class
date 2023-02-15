// Solution to the Exercize 1

// To build and execute (with a sample file) this program run the following 2 commands:
// go build -o ./bin/mostFrequentWord_Ex_1-sol ./src/not-so-simple-program/exercise-1-solution
// ./bin/mostFrequentWord_Ex_1-sol testdata/count-words/file-with-two-unique-words.txt

package main

import (
	"fmt"

	"github.com/EnricoPicci/go-class/src/not-so-simple-program/helpers"
)

func main() {
	file := helpers.ReadFirstCmdLineArg()

	mostFrequent := mostFrequentWordWithCount(file)

	fmt.Printf("The most frequent word is ==>> \"%v\" \n", mostFrequent.Word)
	fmt.Printf("It is found %v times in the file %v \n", mostFrequent.Occurrencies, file)
}
