package examples_test

import (
	"fmt"

	"github.com/EnricoPicci/go-class/src/testing/examples"
)

// ExampleMyApi shows an example of howto write an example that can be executed as a test and is actually
// executed by the go test tool
func ExampleMyApi() {
	input := "A string"

	result := examples.MyApi(input)

	fmt.Println(result)
	// Output:
	// 8
}

// If the following ExampleMyApiabc is uncommented, then we we that the  Go linter signals an issue since
// MyApiabc is not an exported function (the error message is "ExampleMyApiabc refers to unknown identifier")
// func ExampleMyApiabc() {
// }

// ExampleMyType_Who is an example of example of an exported method of an exported type
func ExampleMyType_Who() {
	aMyType := examples.MyType{}

	result := aMyType.Who()

	fmt.Println(result)
	// Output: I am a value of type MyType

}

// If the following ExampleMyApi_error is uncommented, we see that running go test on this example will generate an
// error as if a real test had failed
// Run the following command to see the error occurring once the example has been uncommented
// go test -timeout 30s -run ^ExampleMyApi_error$ github.com/EnricoPicci/go-class/src/testing/examples -count=1
// func ExampleMyApi_error() {
// 	panic("example failing")
// 	// Output: I am a value of type MyType
// }
