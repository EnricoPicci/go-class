package local

import (
	"os"
)

func NewFile(fName string) (*os.File, error) {
	return os.Create(fName)
}
