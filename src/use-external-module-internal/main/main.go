package main

import "github.com/EnricoPicci/go-class-hello-with-internal/src/pkg/hello"

//
// the following import rises an error since we are trying to use a package that is considered internal since it resides below an "internal" folder
// import "github.com/EnricoPicci/go-class-hello-with-internal/src/internal/print"

func main() {
	hello.Upper("Ciao world")
}
