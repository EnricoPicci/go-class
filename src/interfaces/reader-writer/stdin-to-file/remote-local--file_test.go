package main

import (
	"io"
	"os"
	"path"
	"testing"

	"github.com/EnricoPicci/go-class/src/interfaces/reader-writer/stdin-to-file/testhelpers"
	"github.com/EnricoPicci/go-class/src/interfaces/reader-writer/stdin-to-file/writers/local"
	"github.com/EnricoPicci/go-class/src/interfaces/reader-writer/stdin-to-file/writers/remote"
	helpers "github.com/EnricoPicci/go-class/src/testhelpers"
)

// test
func TestClient(t *testing.T) {

	quitCmd := "quit"

	strings := []string{
		"abc\n", "123456\n", "", "00\n", quitCmd + "\n",
	}
	expectedNumBytesCopied := int64(14)
	var numBytesCopied int64

	var sr *testhelpers.StringsReader
	var stdin *stdin

	// =========  test copy to remote file  =========
	//
	sr = testhelpers.NewStringsReader(strings)
	stdin = newStdin(quitCmd, sr)
	// Create the remote file
	remoteFileName := "remote.txt"
	remoteF, resp, err := remote.NewFile(remoteFileName)
	if err != nil {
		t.Log(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		t.Errorf("Error received %v after remote.NewFile\n", err)
		handleRemoteError(err, remoteFileName, resp)
	}
	// copy from stdin
	numBytesCopied, err = io.Copy(remoteF, stdin)
	if numBytesCopied != expectedNumBytesCopied {
		t.Fatalf("Number of bytes copied to remote %v - expected %v\n", numBytesCopied, expectedNumBytesCopied)
	}
	if err != nil {
		t.Fatalf("Error received while running copyFromStdin to remote file:  %v\n", err)
	}
	// close the remote file
	err = remoteF.Close()
	if err != nil {
		t.Fatalf("Error received while closing the remote file %v:  %v\n", remoteF.Name, err)
	}

	// =========  test copy to local file  =========
	//
	sr = testhelpers.NewStringsReader(strings)
	stdin = newStdin(quitCmd, sr)
	// Create the local file
	localFileName := "local.txt"
	prjDir := helpers.ProjectDir()
	localFilePath := path.Join(prjDir, "out", localFileName)
	localF, err := local.NewFile(localFilePath)
	if err != nil {
		t.Errorf("Error received %v after NewFile for local file\n", err)
	}
	// copy from stdin
	numBytesCopied, err = io.Copy(localF, stdin)
	if numBytesCopied != expectedNumBytesCopied {
		t.Errorf("Number of bytes copied to local %v - expected %v\n", numBytesCopied, expectedNumBytesCopied)
	}
	if err != nil {
		t.Errorf("Error received while running copyFromStdin to local file:  %v\n", err)
	}
	// close
	err = localF.Close()
	if err != nil {
		t.Errorf("Error received while closing the local file %v:  %v\n", localF.Name(), err)
	}

	// check that the files copied to local and remote file are the same
	// the test assumes that the server runs locally and writes on the ./out folder
	fileCopiedToRemote := path.Join(prjDir, "out", remoteFileName) // the test assumes that the server runs locally and writes on the ./out folder
	fileCopiedToLocal := path.Join(prjDir, "out", localFileName)
	bRemote, err := os.ReadFile(fileCopiedToRemote)
	if err != nil {
		t.Errorf("Error received while opening the remote file %v:  %v\n", fileCopiedToRemote, err)
	}
	bLocal, err := os.ReadFile(fileCopiedToLocal)
	if err != nil {
		t.Errorf("Error received while opening the local file %v:  %v\n", fileCopiedToLocal, err)
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
