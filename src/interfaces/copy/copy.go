// In this example we use the io.Copy function to read from stdin and write to a file.
// stdin is a value of type *os.File which implements the io.Reader interface
// a file is value of type *os.File which implements the io.Writer interface
// We wrap the file in a type defined in this package to be able to quit the program programmatically.
//
// To run the program use the following command
// go run ./src/interfaces/copy

package main

import (
	"fmt"
	"io"
	"os"
)

type file struct {
	file *os.File
}

func (f *file) Write(b []byte) (int, error) {
	if string(b) == "quit\n" {
		os.Exit(0)
	}
	i, err := f.file.Write(b)
	return i, err
}

func main() {
	fName := "./out/copy.txt"

	fmt.Printf("Enter what you want to write in the %v file (exit typing 'quit')\n", fName)

	f, err := os.Create(fName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	dst := &file{f}

	_, err = io.Copy(dst, os.Stdin)

	if err != nil {
		fmt.Println(err)
	}

}
