package main

import (
	"runtime"
	"testing"
)

var result int

func BenchmarkAddLineNumbersToFilesInDir(b *testing.B) {
	const howMany = 8
	const concurrent = 2

	// runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	var r int
	for i := 0; i < b.N; i++ {
		r += callServer(howMany, concurrent)
	}
	result = r
	// 12647
}
