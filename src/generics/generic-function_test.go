package generics

import (
	"testing"
)

func TestIdentity(t *testing.T) {
	inString := "I am a string"
	inInt := 123
	inUserDefType := struct{ name string }{"I am the name"}

	// test the identity function
	var expected any
	var obtained any

	expected = inString
	obtained = identity(inString)
	checkResult(t, expected, obtained)

	expected = inInt
	obtained = identity(inInt)
	checkResult(t, expected, obtained)

	expected = inUserDefType
	obtained = identity(inUserDefType)
	checkResult(t, expected, obtained)
}

func checkResult(t *testing.T, expected any, obtained any) {
	if expected != obtained {
		t.Errorf("expected %v, got %v", expected, obtained)
	}
}
