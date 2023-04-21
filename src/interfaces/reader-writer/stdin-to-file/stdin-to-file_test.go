package main

import (
	"os"
	"path"
	"testing"

	"github.com/EnricoPicci/go-class/src/interfaces/reader-writer/stdin-to-file/testhelpers"
	helpers "github.com/EnricoPicci/go-class/src/testhelpers"
)

func TestCcopyFromStdinTo(t *testing.T) {

	quitCmd := "quit"

	strings := []string{
		"abc\n", "123456\n", "", "00\n", quitCmd + "\n",
	}

	var sr *testhelpers.StringsReader
	var stdin *stdin

	// =========  test copy to remote file  =========
	//
	sr = testhelpers.NewStringsReader(strings)
	stdin = newStdin(quitCmd, sr)
	remoteFileName := "remote.txt"

	// copy from stdin
	copyFromStdinTo(remoteFileName, true, stdin)

	// =========  test copy to local file  =========
	//
	sr = testhelpers.NewStringsReader(strings)
	stdin = newStdin(quitCmd, sr)
	localFileName := "local.txt"
	prjDir := helpers.ProjectDir()
	localFilePath := path.Join(prjDir, "out", localFileName)
	// copy from stdin
	copyFromStdinTo(localFilePath, false, stdin)

	// check that the files copied to local and remote file are the same
	// the test assumes that the server runs locally and writes on the ./out folder
	fileCopiedToRemote := path.Join(prjDir, "out", remoteFileName) // the test assumes that the server runs locally and writes on the ./out folder
	bRemote, err := os.ReadFile(fileCopiedToRemote)
	if err != nil {
		t.Errorf("Error received while opening the remote file %v:  %v\n", fileCopiedToRemote, err)
	}
	bLocal, err := os.ReadFile(localFilePath)
	if err != nil {
		t.Errorf("Error received while opening the local file %v:  %v\n", localFilePath, err)
	}
	if string(bRemote) != string(bLocal) {
		t.Error("The content copied to remote file is not equal to the content copied to local file\n")
		t.Errorf("The content copied to remote file: %v\n", string(bRemote))
		t.Errorf("The content copied to local  file: %v\n", string(bLocal))
	}
	// check that the copied content is equal to the stdin input
	var stdinInput string
	for i := 0; i < len(strings)-1; i++ {
		stdinInput = stdinInput + strings[i]
	}
	if string(bRemote) != stdinInput {
		t.Error("The content copied is not equal to the content of the stdin\n")
		t.Errorf("The content copied: %v\n", string(bRemote))
		t.Errorf("The content input : %v\n", stdinInput)
	}
}
