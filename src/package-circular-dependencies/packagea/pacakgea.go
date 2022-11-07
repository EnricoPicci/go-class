package packagea

import "github.com/EnricoPicci/go-class/src/package-circular-dependencies/packageb"

func DoStuffA() {
	println("I am Package A")
	packageb.DoStuffB()
}
