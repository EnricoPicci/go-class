package testhelpers

import (
	"errors"
	"fmt"
)

type StringsReader struct {
	i       int
	strings []string
}

func NewStringsReader(strings []string) *StringsReader {
	return &StringsReader{0, strings}
}

func (sr *StringsReader) Read(b []byte) (int, error) {
	if sr.i == len(sr.strings) {
		return 0, errors.New("EOF")
	}
	s := sr.strings[sr.i]
	if len(b) < len(s) {
		msg := fmt.Sprintf("byte array size %v smaller than string %v", len(b), s)
		panic(msg)
	}
	for j := 0; j < len(s); j++ {
		b[j] = s[j]
	}
	sr.i++
	return len(s), nil
}
