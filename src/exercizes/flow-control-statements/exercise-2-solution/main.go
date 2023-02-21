package main

import (
	"fmt"

	exercizehelpers "github.com/EnricoPicci/go-class/src/exercize-helpers"
)

func main() {
	words := exercizehelpers.ReadCmdLineArgs()

	fmt.Println("These are the words you have entered")
	fmt.Println(words)

	fmt.Println("Here the check results of the words you have entered")

	checkManyWords(words)
}
