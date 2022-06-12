package testhelpers

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

// this is the way to get the current working directory even if the test is launched from the vscode debugger
// it is the same as pwd := os.Getenv("PWD") which works only if the test is launched from the command line
// https://stackoverflow.com/a/38644571/5699993
func FilePath(fileName string) string {
	prjDir := "go-class"

	var (
		_, b, _, _ = runtime.Caller(0)
		pwd        = filepath.Dir(b)
	)

	prjDirSplit := strings.Split(pwd, "go-class")
	if len(prjDirSplit) != 2 {
		panic(fmt.Sprintf("The name of the project directory (%v) is not in the program working directory pwd (%v)\n", "go-class", pwd))
	}

	return prjDirSplit[0] + prjDir + "/" + fileName
}
