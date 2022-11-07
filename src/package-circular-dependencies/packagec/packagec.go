package packagec

// import "github.com/EnricoPicci/go-class/src/package-circular-dependencies/packagea" <==== uncomment this line

func DoStuffC() {
	println("I am Package C")
	//  packagea.DoStuffA()  <==== uncomment this line
}
