package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var data []string

// tracks the number of concurrent reads
var readCount int64

var rwMu sync.RWMutex

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	// writer goroutine
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			writer(i)
		}
		wg.Done()
	}()

	// reader goroutines
	for i := 0; i < 8; i++ {
		id := i
		go func() {
			for {
				reader(id)
			}
		}()
	}

	wg.Wait()

}

func writer(i int) {
	rwMu.Lock()
	rc := atomic.LoadInt64(&readCount)
	fmt.Printf(">>>>>>> write %d, readCount %d\n", i, rc)
	data = append(data, fmt.Sprintf("String: %d", i))
	rwMu.Unlock()
}

func reader(id int) {
	rwMu.RLock()
	rc := atomic.AddInt64(&readCount, 1)
	l := len(data)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	fmt.Printf("Goroutine id %v --- Reading >>>>>> readCount %d -- length %d\n", id, rc, l)
	atomic.AddInt64(&readCount, -1)
	rwMu.RUnlock()
}
