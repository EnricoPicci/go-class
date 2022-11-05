package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/EnricoPicci/go-class/src/interfaces/reader-writer/stdin-to-file/writers/local"
	"github.com/EnricoPicci/go-class/src/interfaces/reader-writer/stdin-to-file/writers/remote"
)

func CopyFromStdinTo(fileName string, copyToRemote bool, quitCmd string) {
	var localRemote string
	switch copyToRemote {
	case true:
		localRemote = "remote"
	case false:
		localRemote = "local"
	}

	fmt.Printf("Enter what you want to write to the %v file %v (exit typing %v)\n", localRemote, fileName, quitCmd)

	_quitCmd := quitCmd + "\n"
	stdin := newStdin(_quitCmd, os.Stdin)

	copyFromStdinTo(fileName, copyToRemote, stdin)
}

func copyFromStdinTo(fileName string, copyToRemote bool, src *stdin) {
	var file io.WriteCloser
	switch copyToRemote {
	case true:
		var err error
		var resp string
		file, resp, err = remote.NewFile(fileName)
		if err != nil {
			handleRemoteError(err, fileName, resp)
		}
	case false:
		var err error
		file, err = local.NewFile(fileName)
		if err != nil {
			fmt.Printf("Failed opening local file %v\n", fileName)
			log.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}

	n, err := io.Copy(file, src)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Copied %v bytes\n", n)

	// close the remote file
	err = file.Close()
	if err != nil {
		log.Fatalf("Error received while closing the remote file %v:  %v\n", fileName, err)
	}
}

func handleRemoteError(err error, fileName string, resp string) {

	if err == remote.NoServerErr {
		fmt.Println("No server available")
		fmt.Println("You may need to start the server")
		os.Exit(1)
	}
	if err == remote.TimeoutErr {
		fmt.Println("Timeout while trying to reach the server")
		fmt.Println("Check the status of the server")
		os.Exit(1)
	}
	fmt.Printf("Failed opening remote file %v\n", fileName)
	log.Printf("Message reveived from the server: %v\n", resp)
	log.Printf("Error: %v\n", err)
	os.Exit(1)
}
