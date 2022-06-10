package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestAddLineNumber(t *testing.T) {
	prjDir := "go-class"

	// this is the way to get the current working directory even if the test is launched from the vscode debugger
	// it is the same as pwd := os.Getenv("PWD") which works only if the test is launched from the command line
	// https://stackoverflow.com/a/38644571/5699993
	var (
		_, b, _, _ = runtime.Caller(0)
		pwd        = filepath.Dir(b)
	)

	prjDirSplit := strings.Split(pwd, prjDir)
	if len((prjDirSplit)) != 2 {
		panic(fmt.Sprintf("The name of the project directory (%v) is not in the program working directory pwd (%v)\n", prjDir, pwd))
	}

	var path = prjDirSplit[0] + prjDir + "/canti-divina-commedia/01 - Inferno - CANTO PRIMO.txt"
	numberedLines := addLineNumber(path)

	// test the result
	numberOfLines := len(numberedLines)
	expected := 137
	if numberOfLines != expected {
		t.Errorf("The result is %v instead of %v", numberOfLines, expected)
	}
}
