package testhelpers

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// this is the way to get this project directory even from tests launched from the vscode debugger
// it is the same as pwd := os.Getenv("PWD") which works only if the test is launched from the command line
// https://stackoverflow.com/a/38644571/5699993
func ProjectDir() string {
	srcDir := "src"

	var (
		_, b, _, _ = runtime.Caller(0)
		pwd        = filepath.Dir(b)
	)

	prjDirSplit := strings.Split(pwd, srcDir)
	if len(prjDirSplit) > 2 {
		panic(fmt.Sprintf("The name of the source directory (%v) is found more than once in the path of the caller function (%v)\n", srcDir, pwd))
	}
	if len(prjDirSplit) < 2 {
		panic(fmt.Sprintf("The name of the source directory (%v) is not in the path of the caller function (%v)\n", srcDir, pwd))
	}

	return prjDirSplit[0]
}

func FilePath(fileName string) string {
	prjDir := ProjectDir()
	fPath := path.Join(prjDir, fileName)
	return fPath
}
