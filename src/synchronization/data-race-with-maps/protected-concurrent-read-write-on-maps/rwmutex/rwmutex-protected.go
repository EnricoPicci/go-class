package rwmutex

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func ProtectedConcurrentReadWrite(rangeOfKeys int, writes int, readGoroutines int, reads int, delay time.Duration) {
	var wgWrite sync.WaitGroup
	var mu sync.RWMutex

	concurrentlyReadWrittenMap := make(map[int]int)
	readHits := 0

	wgWrite.Add(1)
	go launchWriter(concurrentlyReadWrittenMap, rangeOfKeys, writes, &wgWrite, &mu, delay)

	var wgRead sync.WaitGroup
	wgRead.Add(1)
	go launchReaders(concurrentlyReadWrittenMap, &readHits, rangeOfKeys, reads, readGoroutines, &wgRead, &mu, delay)

	wgWrite.Wait()
	wgRead.Wait()

	log.Printf("RWMutex - written %v keys (in the range 0-%v) - %v read hits (in %v read operations)\n",
		len(concurrentlyReadWrittenMap), rangeOfKeys, readHits, reads)

}

func launchWriter(concurrentlyReadWrittenMap map[int]int, rangeOfKeys int, writes int, wg *sync.WaitGroup, mu *sync.RWMutex, delay time.Duration) {
	defer wg.Done()

	start := time.Now()
	for i := 0; i < writes; i++ {
		rInt := rand.Intn(rangeOfKeys)
		time.Sleep(delay)
		mu.Lock()
		concurrentlyReadWrittenMap[rInt] = rInt
		mu.Unlock()
	}

	elapsed := time.Since(start)
	log.Printf("RWMutex Write finished. Elapsed: %s", elapsed)
}

func launchReaders(concurrentlyReadWrittenMap map[int]int, readHits *int, rangeOfKeys int, reads int, readGoroutines int, wg *sync.WaitGroup, mu *sync.RWMutex, delay time.Duration) {
	defer wg.Done()

	start := time.Now()
	var wgReads sync.WaitGroup
	wgReads.Add(readGoroutines)
	for i := 0; i < readGoroutines; i++ {
		go protectedMapReader(concurrentlyReadWrittenMap, readHits, rangeOfKeys, mu, reads, &wgReads, delay)
	}
	wgReads.Wait()

	elapsed := time.Since(start)
	log.Printf("RWMutex Read finished. Elapsed: %s", elapsed)
}

func protectedMapReader(concurrentlyReadWrittenMap map[int]int, readHits *int, rangeOfKeys int, mu *sync.RWMutex, reads int, wg *sync.WaitGroup, delay time.Duration) {
	defer wg.Done()

	for i := 0; i < reads; i++ {
		rInt := rand.Intn(rangeOfKeys)
		time.Sleep(delay)
		mu.RLock()
		_, found := concurrentlyReadWrittenMap[rInt]
		mu.RUnlock()
		if found {
			*readHits++
		}
	}
}
