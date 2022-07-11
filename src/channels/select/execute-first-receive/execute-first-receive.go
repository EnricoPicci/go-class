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
	go sendToChannelWithDelay(ch_2, "xyz")

	var valReceived string

	select {
	case valReceived = <-ch_1:
		fmt.Println("Recevied on channel ch_1")
		fmt.Printf("Value %v\n", valReceived)
	case valReceived = <-ch_2:
		fmt.Println("Recevied on channel ch_2")
		fmt.Printf("Value %v\n", valReceived)
	}

	time.Sleep(2000)

}

func sendToChannelWithDelay(ch chan string, val string) {
	// random delay
	delay := rand.Intn(1000000)
	fmt.Printf("delay %v in sending %v\n", delay, val)
	time.Sleep(time.Duration(delay))

	ch <- val
}
