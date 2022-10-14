// go run ./src/exercizes/400-built-in-type-conversions

package main

import "fmt"

func main() {
	var startingNumber float32 = 100
	var two int = 2

	multiplicationBy2 := startingNumber * 2
	fmt.Printf("Multiply %v by 2: %v - the result is a value of type %T \n", startingNumber, multiplicationBy2, multiplicationBy2)

	// the following line does not compile since the types of the 2 operands are different: float32 and int
	// multiplyAgain := startingNumber * two

	// this line does compile since we convert the int value to float32
	multiplyAgainRight := startingNumber * float32(two)
	fmt.Printf("Convert to float32 and then multiply %v by %v: %v - the result is a value of type %T  \n", startingNumber, two, multiplyAgainRight, multiplyAgainRight)
}
