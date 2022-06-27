package main

import (
	"fmt"
	"sync"
)

var sendReceive = make(map[int]string)
var mu sync.Mutex

func main() {

	for i := 0; i < 100000; i++ {
		WaitForTask(i)
	}

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

func WaitForTask(i int) {
	ch := make(chan string)

	go func() {
		<-ch
		mu.Lock()
		_, found := sendReceive[i]
		if !found {
			sendReceive[i] = "R"
		}
		mu.Unlock()
	}()

	ch <- "stuff"
	mu.Lock()
	_, found := sendReceive[i]
	if !found {
		sendReceive[i] = "S"
	}
	mu.Unlock()
}
