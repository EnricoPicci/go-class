// Run the program with the command
// go run ./src/file-system/read-file/.

package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func main() {
	// thisFile is the file path of this source file you are now reading
	var _, thisFile, _, _ = runtime.Caller(0)
	printContent(thisFile)

	aNonExistingFile := "fileThatDoesNotExist.gooo"
	printContent(aNonExistingFile)
}

func printContent(file string) {
	dat, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Error encountered reading %v - Error: %v", file, err.Error())
	}
	fmt.Printf("This is the content of file %v \n", file)
	fmt.Println(string(dat))
}
