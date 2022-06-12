package testhelpers

import (
	"os"
	"testing"
)

// this test can not be run from the vscode debugger since the logic of the testhelpers FilePath function
// is implemented to get the current working directory even if the test is launched from the vscode debugger
func TestAddLineNumber(t *testing.T) {
	fileName := "abc.txt"
	testDirPath := "src/testhelpers/"

	expected := os.Getenv("PWD") + "/" + fileName

	fPath := FilePath(testDirPath + fileName)

	if fPath != expected {
		t.Errorf("The result is %v instead of %v", fPath, expected)
	}

}
