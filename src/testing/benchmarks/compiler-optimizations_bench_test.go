// Example of benchmark that is affected by compiler optimizations
// https://github.com/golang/go/issues/27400#issuecomment-546513538

package benchmarkspkg

import (
	"runtime"
	"testing"
	"time"
)

// BenchmarkX results are misleading since the compiler optimizes and remove the code in the for loop closure
// The other benchmarks, which exercize the same code, show ways of avoiding compiler optimizations
func BenchmarkX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f()
	}
}

var sink int

func BenchmarkXSink(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink = f()
	}
}

func BenchmarkXKeepAlive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(f())
	}
}

var inp float64

func init() {
	if time.Now().Year() > 1900 {
		inp = 123412341234123
	}
}

func f() int {
	x := inp
	x /= 7.3
	x /= 7.3
	x /= 7.3
	x /= 7.3
	return int(x)
}
