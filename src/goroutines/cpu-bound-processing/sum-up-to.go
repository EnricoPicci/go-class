package main

import (
	"math"
)

func sumUpTo(upTo int, concurrent int) int {
	ranges := calcRanges(upTo, concurrent)
	if len(ranges) != concurrent {
		panic("The number of ranges is not equal to the number of concurrent goroutines foreseen")
	}

	var sum int

	results := make(chan int)

	for _, r := range ranges {
		go sumFromTo(r[0], r[1], results)
	}

	for i := 0; i < concurrent; i++ {
		sum += <-results
	}

	return sum
}

func calcInterval(upTo int, concurrent int) int {
	return int(math.Floor(float64(upTo) / float64(concurrent)))
}

func calcRanges(upTo int, concurrent int) [][]int {
	interval := calcInterval(upTo, concurrent)

	start := 1
	end := interval
	ranges := make([][]int, concurrent)
	for i := 0; i < concurrent; i++ {
		ranges[i] = []int{start, end}
		start = end + 1
		end += interval
	}
	if ranges[concurrent-1][1] < upTo {
		ranges[concurrent-1][1] = upTo
	}

	return ranges
}

func sumFromTo(from int, to int, results chan int) {
	var sum int
	for i := from; i <= to; i++ {
		sum += i
	}
	results <- sum
}
