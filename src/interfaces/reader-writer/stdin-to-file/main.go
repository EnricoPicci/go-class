// go run ./src/interfaces/reader-writer/stdin-to-file/ -file test.txt

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	fNamePtr := flag.String("file", "", "Name of the target file")
	remotePtr := flag.Bool("remote", false, "Remote")
	quitCmdPtr := flag.String("quitCmd", "quit", "Command to quit copy")
	flag.Parse()

	if *fNamePtr == "" {
		fmt.Println("Specify the name of the file you want to copy to")
		os.Exit(1)
	}

	CopyFromStdinTo(*fNamePtr, *remotePtr, *quitCmdPtr)
}
