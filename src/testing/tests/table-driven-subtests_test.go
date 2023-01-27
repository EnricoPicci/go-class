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
			if got := FunctionToTest(tc.input, delayInMillisec); got != tc.output {
				t.Errorf("got %v; want %v", got, tc.output)
			}
		})
	}
}
