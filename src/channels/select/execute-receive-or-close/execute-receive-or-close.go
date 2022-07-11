package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ch_1 := make(chan string)
	ch_2 := make(chan string)

	go sendToChannelWithDelay(ch_1, "abc")
	go closeChannelWithDelay(ch_2, "ch_2")

	var valReceived string

	select {
	case valReceived = <-ch_1:
		fmt.Println("Recevied on channel ch_1")
		fmt.Printf("Value %v\n", valReceived)
	case _, chOpen := <-ch_2:
		if chOpen {
			panic("should never arrive here")
		}
		fmt.Println("ch_2 has been closed")
	}

	time.Sleep(300000)

}

func sendToChannelWithDelay(ch chan string, val string) {
	// random delay
	delay := rand.Intn(100000)
	time.Sleep(time.Duration(delay))
	fmt.Printf("delay %v in sending %v\n", delay, val)

	ch <- val
}

func closeChannelWithDelay(ch chan string, chName string) {
	// random delay
	delay := rand.Intn(100000)
	time.Sleep(time.Duration(delay))
	fmt.Printf("delay %v in closing the channel %v\n", delay, chName)

	close(ch)
}
