package packageb

import "github.com/EnricoPicci/go-class/src/package-circular-dependencies/packagec"

func DoStuffB() {
	println("I am Package B")
	packagec.DoStuffC()
}
