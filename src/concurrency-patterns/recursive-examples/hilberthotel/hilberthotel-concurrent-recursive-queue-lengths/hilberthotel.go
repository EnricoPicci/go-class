package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/EnricoPicci/go-class/src/concurrency-patterns/recursive-examples/hilberthotel"
)

func RoomKeysClerk(upTo int, keysCh chan<- int) {
	for i := 0; i < upTo; i++ {
		keysCh <- i + 1
	}
	close(keysCh)
}

type QueueLengths struct {
	lengths   []int
	busNumber int
}

func BusClerk(busNumber int, roomKeysCh <-chan int, welcomeKitsCh chan<- []hilberthotel.WelcomeKit, queueLengthsCh chan<- QueueLengths, parallelism int) {
	delay := 10 * time.Microsecond
	var count = 0
	var passengerNumber = 1
	var nextClerkCh chan int

	queueLengths := QueueLengths{[]int{}, busNumber}

	welcomeKits := []hilberthotel.WelcomeKit{}

	for roomKey := range roomKeysCh {
		queueLengths.lengths = append(queueLengths.lengths, len(roomKeysCh))
		count++
		if nextClerkCh == nil {
			nextClerkCh = make(chan int, parallelism)
			go BusClerk(busNumber+1, nextClerkCh, welcomeKitsCh, queueLengthsCh, parallelism)
		}
		if count == passengerNumber {
			kit := hilberthotel.NewWelcomeKit(busNumber, passengerNumber, roomKey, delay)
			welcomeKits = append(welcomeKits, kit)
			passengerNumber++
			count = 0
			continue
		}
		nextClerkCh <- roomKey
	}

	if nextClerkCh != nil {
		welcomeKitsCh <- welcomeKits
		queueLengthsCh <- queueLengths
		close(nextClerkCh)
	} else {
		close(welcomeKitsCh)
		close(queueLengthsCh)
	}
}

func GoHilbert(upTo int, parallelism int) ([]hilberthotel.WelcomeKit, []QueueLengths) {
	if parallelism < 0 {
		parallelism = 0
	}
	keysCh := make(chan int, parallelism)
	go RoomKeysClerk(upTo, keysCh)

	hilbertCh := make(chan []hilberthotel.WelcomeKit, parallelism)
	queueLengthsCh := make(chan QueueLengths, parallelism)
	go BusClerk(1, keysCh, hilbertCh, queueLengthsCh, parallelism)

	queueLengths := []QueueLengths{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for qL := range queueLengthsCh {
			queueLengths = append(queueLengths, qL)
		}
	}()

	kits := []hilberthotel.WelcomeKit{}
	for busKits := range hilbertCh {
		kits = append(kits, busKits...)
	}

	wg.Wait()

	fmt.Println()
	fmt.Printf("%v guests have been given a room by Hilber at his Hotel\n", len(kits))

	return kits, queueLengths
}
