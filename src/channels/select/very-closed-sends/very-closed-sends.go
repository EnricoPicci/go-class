package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	runtime.GOMAXPROCS(1)

	ch_1 := make(chan string)
	ch_2 := make(chan string)

	firstDelay := 100
	delayDifference := 100000

	go sendToChannelWithDelay(ch_1, "abc", firstDelay)
	go sendToChannelWithDelay(ch_2, "xyz", firstDelay+delayDifference)

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

func sendToChannelWithDelay(ch chan string, val string, delay int) {
	time.Sleep(time.Duration(delay))

	ch <- val
}
