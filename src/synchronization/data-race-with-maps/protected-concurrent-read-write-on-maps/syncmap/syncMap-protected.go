package syncmap

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func ProtectedConcurrentReadWrite(rangeOfKeys int, writes int, readGoroutines int, reads int, delay time.Duration) {
	var wgWrite sync.WaitGroup
	wgWrite.Add(1)

	var syncMap sync.Map
	readHits := 0

	go launchWriter(&syncMap, rangeOfKeys, writes, &wgWrite, delay)

	var wgRead sync.WaitGroup
	wgRead.Add(1)
	go launchReaders(&syncMap, &readHits, rangeOfKeys, reads, readGoroutines, &wgRead, delay)

	wgWrite.Wait()
	wgRead.Wait()

	log.Printf("Sync Map - written %v keys (in the range 0-%v) - %v read hits (in %v read operations)\n",
		lenSyncMap(&syncMap), rangeOfKeys, readHits, reads)

}

func launchWriter(concurrentlyReadWrittenMap *sync.Map, rangeOfKeys int, writes int, wg *sync.WaitGroup, delay time.Duration) {
	defer wg.Done()

	start := time.Now()
	for i := 0; i < writes; i++ {
		rInt := rand.Intn(rangeOfKeys)
		time.Sleep(delay)
		concurrentlyReadWrittenMap.Store(rInt, rInt)
	}

	elapsed := time.Since(start)
	log.Printf("Sync Map Write finished. Elapsed: %s", elapsed)
}

func launchReaders(concurrentlyReadWrittenMap *sync.Map, readHits *int, rangeOfKeys int, reads int, readGoroutines int, wg *sync.WaitGroup, delay time.Duration) {
	defer wg.Done()

	start := time.Now()
	var wgReads sync.WaitGroup
	wgReads.Add(readGoroutines)

	for i := 0; i < readGoroutines; i++ {
		go protectedMapReader(concurrentlyReadWrittenMap, readHits, rangeOfKeys, reads, &wgReads, delay)
	}
	wgReads.Wait()

	elapsed := time.Since(start)
	log.Printf("Sync Map Read finished. Elapsed: %s", elapsed)
}

func protectedMapReader(concurrentlyReadWrittenMap *sync.Map, readHits *int, rangeOfKeys int, reads int, wg *sync.WaitGroup, delay time.Duration) {
	defer wg.Done()

	for i := 0; i < reads; i++ {
		rInt := rand.Intn(rangeOfKeys)
		time.Sleep(delay)
		_, found := concurrentlyReadWrittenMap.Load(rInt)
		if found {
			*readHits++
		}
	}
}
func lenSyncMap(m *sync.Map) int {
	var i int
	m.Range(func(k, v interface{}) bool {
		i++
		return true
	})
	return i
}
