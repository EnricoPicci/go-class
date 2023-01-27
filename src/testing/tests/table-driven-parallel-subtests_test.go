package tests

import (
	"testing"
)

// Table driven subtests give the possibility to exercize the same code with diferent input.
// Subtests can be run selectively since each subtest has a name and the go test tool supports selection of tests via the -run flag
// Subtests can also be run in parallel.

// All the subtests of a table driven test function can be run in parallel if we call t.Parallel() within the subtest function
// https://go.dev/blog/subtests#run-a-group-of-tests-in-parallel
func Test_Subtests_Parallel(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output int
	}{
		{"firstTest", "1234", 4},
		{"secondTest", "abcdef", 6},
		{"thirdTest", "", 0},
	}

	const delayInMillisec = 500

	for _, tc := range testCases {
		// it is important to capture the variable since the function representing the subtest is run in its own goroutine
		// and therefore there can be many goroutines running concurrently and we want that the scope of the _tc variable
		// is the inner closure.
		// On the contrary the scope of the tc variable defined in the loop is outside the closure of the loop and therefore
		// there the tc variable is the same in all runs of the loop and would be shared by all the goroutines running the
		// subtests unless we capture the range variable value in _tc variable.
		// See this
		// https://github.com/golang/go/wiki/CommonMistakes#using-reference-to-loop-iterator-variable
		_tc := tc // capture range variable
		t.Run(_tc.name, func(t *testing.T) {
			t.Parallel()
			if got := FunctionToTest(_tc.input, delayInMillisec); got != _tc.output {
				t.Errorf("got %v; want %v", got, _tc.output)
			}
		})
	}
}
