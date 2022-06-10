package main

import (
	"runtime"
	"testing"
)

var result int

func BenchmarkSumUpTo_1000000(b *testing.B) {
	const upTo = 1000000
	var concurrent = runtime.NumCPU()
	//const concurrent = 1
	runtime.GOMAXPROCS(runtime.NumCPU())
	// runtime.GOMAXPROCS(1)
	var r int
	for i := 0; i < b.N; i++ {
		r = sumUpTo(upTo, concurrent)
	}
	result = r
}
