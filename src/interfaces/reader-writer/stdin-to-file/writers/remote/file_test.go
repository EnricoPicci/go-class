package remote_test

import (
	"strings"
	"testing"
	"time"

	"github.com/EnricoPicci/go-class/src/interfaces/reader-writer/stdin-to-file/writers/remote"
)

// test
func TestNewFile(t *testing.T) {
	fileName := "abc.txt"
	f, r, err := remote.NewFile(fileName)

	if err != nil {
		t.Fatalf("Error received %v after NewFile\n", err)
	}
	if f == nil {
		t.Fatalf("File %v not created\n", fileName)
	}
	t.Log(f)
	substringExpeted := fileName + " created"
	if !strings.Contains(r, substringExpeted) {
		t.Fatalf("The response '%v' does not contain the expected substring '%v'", r, substringExpeted)
	}
	t.Log("Response received:", r)

	testString := "a test string " + time.Now().String()

	n, err := f.Write([]byte(testString))
	if err != nil {
		t.Fatalf("Error received %v after Write\n", err)
	}
	if n != len(testString) {
		t.Fatalf("Expected '%v' - received '%v'", len(testString), n)
	}

	err = f.Close()
	if err != nil {
		t.Fatalf("Error received %v after closing\n", err)
	}

}
