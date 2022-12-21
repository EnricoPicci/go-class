package main

import (
	"fmt"
	"sync"
)

func main() {
	sendFirstCounter := 0
	receiveFirstCounter := 0

	var mu sync.Mutex
	sendReceive := make(map[int]string)

	for i := 0; i < 100000; i++ {
		var ch chan int
		var wg sync.WaitGroup
		wg.Add(2)
		ch = make(chan int)
		// this goroutine RECEIVES data
		go func() {
			<-ch
			mu.Lock()
			receiveFirstCounter++
			_, found := sendReceive[i]
			if !found {
				sendReceive[i] = "R"
			}
			mu.Unlock()
			wg.Done()
		}()

		// this goroutine SENDS data
		go func() {
			ch <- 123
			mu.Lock()
			sendFirstCounter++
			_, found := sendReceive[i]
			if !found {
				sendReceive[i] = "S"
			}
			mu.Unlock()
			wg.Done()
		}()

		wg.Wait()
	}

	fmt.Println(receiveFirstCounter)
	fmt.Println(sendFirstCounter)
	howManyReceivesFirst := 0
	howManySendsFirst := 0
	for _, v := range sendReceive {
		if v == "R" {
			howManyReceivesFirst++
		} else {
			howManySendsFirst++
		}
	}
	fmt.Printf("%v times the Receive goroutine has executed the next line after the send/receive atomic operation before the Send goroutine\n",
		howManyReceivesFirst)
	fmt.Printf("%v times the Send goroutine has executed the next line after the send/receive atomic operation before the Receive goroutine\n",
		howManySendsFirst)
}
