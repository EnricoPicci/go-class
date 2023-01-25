package tests

import (
	"testing"
)

// Table driven subtests give the possibility to exercize the same code with diferent input.
// Subtests can be run selectively since each subtest has a name and the go test tool supports selection of tests via the -run flag
// Subtests can also be run in parallel.

// It is possible to create a series of tests for the same function that exercize the function with different input.
// These tests are typically driven by a table (slice) which the test function loops through.
// The table (slice) typically has a name for the subtest, the input and the expected output.
// In the loop we call a subtest function passing it the name.
// It is possible to run selective subtests using the go test command. For instance to run the "empty_test" subtest
// we can run the command
// go test -run Test_TableDriven_Subtests/empty_test ./src/testing/tests/... -count 1 -v
func Test_TableDriven_Subtests(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output int
	}{
		{"1234_test", "1234", 4},
		{"abcdef_test", "abcdef", 6},
		{"empty_test", "", 0},
	}

	const delayInMillisec = 500

	for _, tc := range testCases {
		// In this case capturing the range variable is not important since each subtest runs sequentially
		// Capturing te range variable becomes important when the subtests run in parallel (see the test with subtests in
		// parallel for more details)
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := FunctionToTest(tc.input, delayInMillisec); got != tc.output {
				t.Errorf("got %v; want %v", got, tc.output)
			}
		})
	}
}

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
