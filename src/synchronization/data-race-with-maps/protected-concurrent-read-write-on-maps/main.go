package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/EnricoPicci/go-class/src/synchronization/data-race-with-maps/protected-concurrent-read-write-on-maps/mutex"
	"github.com/EnricoPicci/go-class/src/synchronization/data-race-with-maps/protected-concurrent-read-write-on-maps/rwmutex"
	"github.com/EnricoPicci/go-class/src/synchronization/data-race-with-maps/protected-concurrent-read-write-on-maps/syncmap"
)

func main() {
	var rangeOfKeys = flag.Int("rangeOfKeys", 1000, "keys of the map are in the range from 0 to rangeOfKeys")
	var writes = flag.Int("writes", 1000, "number of writes to be performed by the writer")
	var reads = flag.Int("reads", 1000, "number of reads to be performed by each reader")
	var readers = flag.Int("readers", 10, "number of readers (each reder is a goroutine)")
	var delayMilliseconds = flag.Int("delay", 1, "each read and write is distanciated by this delay in milliseconds")

	fmt.Println("Run the test with the following parameters")
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("%s: %s - %s\n", f.Name, f.Value, f.Usage)
	})
	fmt.Print("\n\n")

	delay := time.Duration(*delayMilliseconds) * time.Millisecond

	var start time.Time
	var elapsed time.Duration

	// Read write with Mutex
	start = time.Now()

	mutex.ProtectedConcurrentReadWrite(*rangeOfKeys, *writes, *readers, *reads, delay)

	elapsed = time.Since(start)
	fmt.Printf("Mutex: %v writes took %s\n\n", *writes, elapsed)

	// Read write with RWMutex
	start = time.Now()

	rwmutex.ProtectedConcurrentReadWrite(*rangeOfKeys, *writes, *readers, *reads, delay)

	elapsed = time.Since(start)
	fmt.Printf("RWMutex: %v writes took %s\n\n", *writes, elapsed)

	// Read write with sync.Map
	start = time.Now()

	syncmap.ProtectedConcurrentReadWrite(*rangeOfKeys, *writes, *readers, *reads, delay)

	elapsed = time.Since(start)
	fmt.Printf("Sync Map: %v writes took %s\n\n", *writes, elapsed)
}
