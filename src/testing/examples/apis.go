package examples

// MyApi is a very intelligent API that takes a string and returns the length of the string
// It is used to illustrate how to create an example
func MyApi(input string) int {
	return len(input)
}

// MyType is a type that provides an exported method as an API
type MyType struct{}

func (mt MyType) Who() string {
	return "I am a value of type MyType"
}
