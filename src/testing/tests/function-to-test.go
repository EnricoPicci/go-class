package tests

import "time"

// FunctionToTest is a function used in the tests
// this function sleeps a bit before returning the result
func FunctionToTest(in string, sleepMilliseconds int) int {
	delay := time.Millisecond * time.Duration(sleepMilliseconds)
	time.Sleep(delay)
	return len(in)
}
