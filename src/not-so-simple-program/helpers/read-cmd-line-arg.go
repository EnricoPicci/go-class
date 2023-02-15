package helpers

import (
	"log"
	"os"
)

func ReadFirstCmdLineArg() string {
	if len(os.Args) < 2 {
		log.Fatal("No file name passed as argument to the program")
	}
	return os.Args[1]
}
