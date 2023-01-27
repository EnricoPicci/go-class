package benchmarkspkg

import "testing"

var input []string = []string{
	"first",
	"/",
	"second",
	"/",
	"third",
}
var expected = "first/second/third"

func TestConcatWithPlus(t *testing.T) {
	got := ConcatWithPlus(input)
	check(t, got)
}

func TestConcatStringBuilder(t *testing.T) {
	got := ConcatWithStringBuilder(input)
	check(t, got)
}

func TestConcatWithStringBuilderAsWriter(t *testing.T) {
	got := ConcatWithStringBuilderAsWriter(input)
	check(t, got)
}

func check(t *testing.T, got string) {
	if got != expected {
		t.Errorf("Got %v - expected %v", got, expected)
	}
}
