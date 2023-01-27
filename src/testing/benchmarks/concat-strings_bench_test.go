// Benchmark different ways to concatenate strings

package benchmarkspkg

import (
	"fmt"
	"testing"
)

var benchmarks = []struct {
	name        string
	stringSlice []string
}{
	{"10 strings", buildStringSlice(10)},
	{"100 strings", buildStringSlice(100)},
	{"1.000 strings", buildStringSlice(1000)},
}

func Benchmark_Plus(b *testing.B) {
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ConcatWithPlus(benchmark.stringSlice)
			}
		})
	}
}

func Benchmark_Builder(b *testing.B) {
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ConcatWithStringBuilder(benchmark.stringSlice)
			}
		})
	}
}

func Benchmark_BuilderW(b *testing.B) {
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ConcatWithStringBuilderAsWriter(benchmark.stringSlice)
			}
		})
	}
}

func buildStringSlice(size int) []string {
	strSl := make([]string, size)
	for i := 0; i < size; i++ {
		strSl[i] = fmt.Sprintf("String_%v", i)
	}
	return strSl
}
