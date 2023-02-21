package exercizehelpers

import (
	"os"
)

func ReadCmdLineArgs() []string {
	return os.Args[1:]
}
