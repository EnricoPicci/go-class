package main

import "fmt"

func main() {
	var word string

	// first case, word too short
	fmt.Println("Here we should see that the word is too short")
	word = "123"
	checkWord(word)

	// second case, word too long
	fmt.Println("Here we should see that the word is too long")
	word = "12345678901234567890"
	checkWord(word)

	// third case, word just right
	fmt.Println("Here we should see that the word is just right")
	word = "1234567890"
	checkWord(word)
}
