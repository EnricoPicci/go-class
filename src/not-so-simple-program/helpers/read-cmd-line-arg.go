package helpers

import "os"

func ReadFirstCmdLineArg() string {
	return os.Args[1]
}
