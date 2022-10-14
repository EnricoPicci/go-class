// go run ./src/exercizes/400-triangle-area
package main

import "fmt"

type Triangle struct {
	Heigth int
	Width  int
}

func main() {
	tr := Triangle{1, 3}

	area := (tr.Heigth * tr.Width) / 2
	fmt.Printf("The area of the triangle with heigth %v and width %v is %v \n", tr.Heigth, tr.Width, area)

	correctArea := (float64(tr.Heigth) * float64(tr.Width)) / 2
	fmt.Printf("The correct area of the triangle with heigth %v and width %v is %v \n", tr.Heigth, tr.Width, correctArea)
}
