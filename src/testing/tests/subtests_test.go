package tests_test

import (
	"testing"

	"github.com/EnricoPicci/go-class/src/testing/tests"
)

func Test_With_Subtests(t *testing.T) {
	// a subtest
	t.Run("subtest-name", func(t *testing.T) {
		got := tests.FunctionToTest("abc", 0)
		expected := 3
		if got != expected {
			t.Errorf("Got %v, expected %v", got, expected)
		}

		// a sub sub test
		t.Run("sub-subtest-name", func(t *testing.T) {
			got := tests.FunctionToTest("abc123", 0)
			expected := 6
			if got != expected {
				t.Errorf("Got %v, expected %v", got, expected)
			}
		})
	})
}
