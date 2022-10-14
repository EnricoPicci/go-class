package main

import (
	"runtime"
	"testing"

	"github.com/EnricoPicci/go-class/src/testhelpers"
)

var result int

func BenchmarkAddLineNumbersToFilesInDir(b *testing.B) {
	const concurrent = 100

	runtime.GOMAXPROCS(runtime.NumCPU())
	// runtime.GOMAXPROCS(1)
	var r int
	for i := 0; i < b.N; i++ {
		var path = testhelpers.FilePath("/canti-divina-commedia")
		var outDirPath = b.TempDir()

		addLineNumbersToFilesInDir(path, outDirPath, concurrent)
	}
	result = r
}
