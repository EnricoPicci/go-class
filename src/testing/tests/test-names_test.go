package tests

import (
	"testing"
)

// Test_Test_And_Subtest_Name checks that the name of a test is the name of the function while the
// name of a subtest is the combination of the name of the parent test and the name of the test passed to the
// t.Run method as first parameter
func Test_Test_And_Subtest_Name(t *testing.T) {
	gotTestName := t.Name()
	expectedTestName := "Test_Test_And_Subtest_Name"
	if gotTestName != expectedTestName {
		t.Errorf("Test name is %v instead of %v", gotTestName, expectedTestName)
	}

	// a subtest
	subTestName := "subtest-name"
	t.Run(subTestName, func(t *testing.T) {
		gotSubTestName := t.Name()
		expectedSubTestName := expectedTestName + "/subtest-name"
		if gotSubTestName != expectedSubTestName {
			t.Errorf("Test name is %v instead of %v", gotSubTestName, expectedSubTestName)
		}

		// a sub sub test
		subSubTestName := "sub-subtest-name"
		t.Run(subSubTestName, func(t *testing.T) {
			gotSubSubTestName := t.Name()
			expectedSubSubTestName := expectedSubTestName + "/sub-subtest-name"
			if gotSubSubTestName != expectedSubSubTestName {
				t.Errorf("Test name is %v instead of %v", gotSubSubTestName, expectedSubSubTestName)
			}
		})
	})

}
