package main

import (
	"fmt"
)

func main() {
	var path = "./canti-divina-commedia"
	var outDirPath = "./tmp"
	var concurrent = 2
	addLineNumbersToFilesInDir(path, outDirPath, concurrent)
	fmt.Println("Program terminated")
}
