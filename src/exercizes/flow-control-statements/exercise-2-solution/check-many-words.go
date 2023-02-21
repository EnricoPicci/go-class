package main

import "fmt"

func checkManyWords(words []string) {
	fmt.Println("===>>> ", words)
	for i := 0; i < len(words); i++ {
		word := words[i]
		switch length := len(word); {
		case length < 8:
			fmt.Printf("\"%v\" is too short \n", word)
		case length > 16:
			fmt.Printf("\"%v\" is too long \n", word)
		default:
			fmt.Printf("\"%v\" is just right \n", word)
		}
	}
}
