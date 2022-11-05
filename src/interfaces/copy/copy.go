// In this example we use the io.Copy function to read from stdin and write to a file.
// stdin is a value of type *os.File which implements the io.Reader interface
// a file is value of type *os.File which implements the io.Writer interface
// We wrap the file in a type defined in this package to be able to quit the program programmatically.
//
// To run the program use the following command
// go run ./src/interfaces/copy

package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
)

type file struct {
	file *os.File
}

func (f *file) Write(b []byte) (int, error) {
	if string(b) == "quit\n" {
		return 0, errors.New("Quitting")
	}
	i, err := f.file.Write(b)
	return i, err
}

func main() {
	fName := path.Join(".", "out", "copy.txt")

	fmt.Printf("Enter what you want to write in the %v file (exit typing 'quit')\n", fName)

	f, err := os.Create(fName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	defer fmt.Println("Closing")

	dst := &file{f}

	_, err = io.Copy(dst, os.Stdin)

	if err != nil {
		fmt.Println(err)
	}

}
