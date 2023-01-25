package benchmarks

import (
	"flag"
	"testing"
	"time"

	"github.com/EnricoPicci/go-class/src/concurrency-patterns/recursive-examples/hilberthotel"
	closurerecursive "github.com/EnricoPicci/go-class/src/concurrency-patterns/recursive-examples/hilberthotel/hilberthotel-closure-recursive"
	concurrencyrecursive "github.com/EnricoPicci/go-class/src/concurrency-patterns/recursive-examples/hilberthotel/hilberthotel-concurrent-recursive"
	nonrecursive "github.com/EnricoPicci/go-class/src/concurrency-patterns/recursive-examples/hilberthotel/hilberthotel-nonrecursive"
)

var numOfGuests int
var delayMicrosec int

var kits []hilberthotel.WelcomeKit

func init() {
	flag.IntVar(&numOfGuests, "numGuests", 1000, "Number of guests that want to stay at Hilbert's Hotel")
	flag.IntVar(&delayMicrosec, "delayMicrosec", 10, "Delay in microsecs to make a welcome kit (simulates work to be done for each guest)")
}

func delay() time.Duration {
	return time.Duration(delayMicrosec) * time.Microsecond
}

func BenchmarkNonRecursive(b *testing.B) {
	var _kits []hilberthotel.WelcomeKit
	for i := 0; i < b.N; i++ {
		_kits = nonrecursive.Hilbert(numOfGuests, delay(), false)
	}
	kits = _kits
}

func BenchmarkClosureRecursive(b *testing.B) {
	var _kits []hilberthotel.WelcomeKit
	for i := 0; i < b.N; i++ {
		_kits = closurerecursive.Hilbert(numOfGuests, delay(), false)
	}
	kits = _kits
}

func BenchmarkConcurrencyRecursiveBuffer0(b *testing.B) {
	buffer := 0

	var _kits []hilberthotel.WelcomeKit
	for i := 0; i < b.N; i++ {
		_kits = concurrencyrecursive.Hilbert(numOfGuests, buffer, delay(), false)
	}
	kits = _kits
}

func BenchmarkConcurrencyRecursiveBuffer100(b *testing.B) {
	buffer := 100

	var _kits []hilberthotel.WelcomeKit
	for i := 0; i < b.N; i++ {
		_kits = concurrencyrecursive.Hilbert(numOfGuests, buffer, delay(), false)
	}
	kits = _kits
}

func BenchmarkConcurrencyRecursiveBuffer1000(b *testing.B) {
	buffer := 1000

	var _kits []hilberthotel.WelcomeKit
	for i := 0; i < b.N; i++ {
		_kits = concurrencyrecursive.Hilbert(numOfGuests, buffer, delay(), false)
	}
	kits = _kits
}
