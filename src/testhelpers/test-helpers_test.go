package testhelpers

import (
	"os"
	"testing"
)

// this test can not be run from the vscode debugger
// to make it work if launched from the vscode debugger we would have to repeat the implementation code which would be ineffective
func TestFilePath(t *testing.T) {
	fileName := "abc.txt"
	testDirPath := "src/testhelpers/"

	expected := os.Getenv("PWD") + "/" + fileName

	fPath := FilePath(testDirPath + fileName)

	if fPath != expected {
		t.Errorf("The result is %v instead of %v", fPath, expected)
	}

}
