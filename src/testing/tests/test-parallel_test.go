package tests

import (
	"testing"
)

// This file contains 3 tests that are run in parallel (these are the tests that call t.Parallel() at the start
// of their execution).
// Each test tests a function which has a delay.
// The total time of execution of all the tests that run in parallel is not the sum of 3 delays but slightly greater than
// such delay since all the tests are run in parallel
//
// To run the test in parallel run the command
// go test -count 1 -run "^(Test_Parallel_1|Test_Parallel_2|Test_Parallel_3)$"  ./src/testing/tests/...
//
// To appreciate the difference you can run 3 other tests, whose content is the same, which do not call t.Parallel()
// at the start of their execution. Such tests therefore run sequentially and the total execution time is the sum
// of the single test execution time. Check the difference in execution time running the following command
// go test -count 1 -run "^(Test_NonParallel_1|Test_NonParallel_2|Test_NonParallel_3)$"  ./src/testing/tests/...

// PARALLEL TESTS

// delayInMillisec is the delay that the function under test will have
// since the execution is in parallel, then if we test all the tests of this file we see that the
// execution time is close to delayInMillisec and not much longer as in the case of sequential execution
const delayInMillisec = 500

func Test_Parallel_1(t *testing.T) {
	t.Parallel()
	inputString := "0123456789"
	gotRes := FunctionToTest(inputString, delayInMillisec)
	expectedRes := 10
	if gotRes != expectedRes {
		t.Errorf("Test name is %v instead of %v", gotRes, expectedRes)
	}
}

func Test_Parallel_2(t *testing.T) {
	t.Parallel()
	inputString := "0123456789"
	gotRes := FunctionToTest(inputString, delayInMillisec)
	expectedRes := 10
	if gotRes != expectedRes {
		t.Errorf("Test name is %v instead of %v", gotRes, expectedRes)
	}
}

func Test_Parallel_3(t *testing.T) {
	t.Parallel()
	inputString := "0123456789"
	gotRes := FunctionToTest(inputString, delayInMillisec)
	expectedRes := 10
	if gotRes != expectedRes {
		t.Errorf("Test name is %v instead of %v", gotRes, expectedRes)
	}
}

// NON PARALLEL TESTS

func Test_NonParallel_1(t *testing.T) {
	inputString := "0123456789"
	gotRes := FunctionToTest(inputString, delayInMillisec)
	expectedRes := 10
	if gotRes != expectedRes {
		t.Errorf("Test name is %v instead of %v", gotRes, expectedRes)
	}
}

func Test_NonParallel_2(t *testing.T) {
	inputString := "0123456789"
	gotRes := FunctionToTest(inputString, delayInMillisec)
	expectedRes := 10
	if gotRes != expectedRes {
		t.Errorf("Test name is %v instead of %v", gotRes, expectedRes)
	}
}

func Test_NonParallel_3(t *testing.T) {
	inputString := "0123456789"
	gotRes := FunctionToTest(inputString, delayInMillisec)
	expectedRes := 10
	if gotRes != expectedRes {
		t.Errorf("Test name is %v instead of %v", gotRes, expectedRes)
	}
}
