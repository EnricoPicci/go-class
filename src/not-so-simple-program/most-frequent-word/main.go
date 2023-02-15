// This program prints the word which is most frequent in a file.
// The file path is passed as first commnad line argument.

// This is an example of a not-so-simple program that shows:
// - how to use functions defined in the same package or in other packages (the helpers package in this case)
// - how to define a function (mostFrequentWord in this case) called by another function (main in this case)
// - how to declare variables
// - how to perform a loop
// The program does not aim to be precise. For instance if the file has 2 words which have the highest count, only one will be printed.

// To build and execute (with a sample file) this program run the following 2 commands:
// go build -o ./bin/mostFrequentWord ./src/not-so-simple-program/most-frequent-word
// ./bin/mostFrequentWord testdata/count-words/file-with-two-unique-words.txt

package main

import (
	"fmt"

	"github.com/EnricoPicci/go-class/src/not-so-simple-program/helpers"
)

func main() {
	file := helpers.ReadFirstCmdLineArg()

	mostFrequent := mostFrequentWord(file)

	fmt.Printf("The most frequent word is ==>> \"%v\" \n", mostFrequent)
}
