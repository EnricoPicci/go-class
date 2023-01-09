package benchmarks

import (
	"testing"

	"github.com/EnricoPicci/go-class/src/concurrency-patterns/recursive-examples/hilberthotel"
	closurerecursive "github.com/EnricoPicci/go-class/src/concurrency-patterns/recursive-examples/hilberthotel/hilberthotel-closure-recursive"
	concurrencyrecursive "github.com/EnricoPicci/go-class/src/concurrency-patterns/recursive-examples/hilberthotel/hilberthotel-concurrent-recursive"
	nonrecursive "github.com/EnricoPicci/go-class/src/concurrency-patterns/recursive-examples/hilberthotel/hilberthotel-nonrecursive"
)

var numOfPassengers = 10000

var kits []hilberthotel.WelcomeKit

func BenchmarkNonRecursive(b *testing.B) {
	var _kits []hilberthotel.WelcomeKit
	for i := 0; i < b.N; i++ {
		_kits = nonrecursive.GoHilbert(numOfPassengers, false)
	}
	kits = _kits
}

func BenchmarkClosureRecursive(b *testing.B) {
	var _kits []hilberthotel.WelcomeKit
	for i := 0; i < b.N; i++ {
		_kits = closurerecursive.GoHilbert(numOfPassengers, false)
	}
	kits = _kits
}

func BenchmarkConcurrencyRecursiveBuffer0(b *testing.B) {
	buffer := 0
	var _kits []hilberthotel.WelcomeKit
	for i := 0; i < b.N; i++ {
		_kits = concurrencyrecursive.GoHilbert(numOfPassengers, buffer, false)
	}
	kits = _kits
}

func BenchmarkConcurrencyRecursiveBuffer100(b *testing.B) {
	buffer := 100
	var _kits []hilberthotel.WelcomeKit
	for i := 0; i < b.N; i++ {
		_kits = concurrencyrecursive.GoHilbert(numOfPassengers, buffer, false)
	}
	kits = _kits
}

func BenchmarkConcurrencyRecursiveBuffer1000(b *testing.B) {
	buffer := 1000
	var _kits []hilberthotel.WelcomeKit
	for i := 0; i < b.N; i++ {
		_kits = concurrencyrecursive.GoHilbert(numOfPassengers, buffer, false)
	}
	kits = _kits
}
