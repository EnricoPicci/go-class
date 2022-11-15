package main

import (
	"math/rand"
	"sync"
)

var sharedMap map[int]int

func main() {
	sharedMap = make(map[int]int)
	valuesFound := []int{}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go mapWriter(&wg)
	for i := 0; i < 10; i++ {
		go mapReader(valuesFound)
	}
	wg.Wait()
}

func mapWriter(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10000000; i++ {
		rInt := rand.Intn(1000000)
		sharedMap[rInt] = rInt
	}
}

func mapReader(valuesFound []int) {
	for {
		val, found := sharedMap[rand.Intn(1000000)]
		if found {
			valuesFound = append(valuesFound, val)
		}
	}
}
