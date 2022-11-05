package testhelpers

import (
	"testing"
)

// test
func TestStringsReader(t *testing.T) {
	strings := []string{
		"abc", "123456", "", "0000000000",
	}

	stringsRead := []string{}

	sr := NewStringsReader(strings)

	b := make([]byte, 1000)
	for {
		n, err := sr.Read(b)
		if err != nil {
			break
		}
		someBytes := b[:n]
		stringsRead = append(stringsRead, string(someBytes))
	}

	if len(stringsRead) != len(strings) {
		t.Errorf("Read %v strings - Expected %v\n", len(stringsRead), len(strings))
	}

	for i, s := range strings {
		if stringsRead[i] != s {
			t.Errorf("At position %v found %v - Expected %v\n", i, stringsRead[i], s)
		}
	}
}
