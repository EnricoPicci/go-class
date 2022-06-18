package main

import (
	"math/rand"
	"sync"
)

var sharedMap map[int]int

func main() {
	sharedMap = make(map[int]int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go mapWriter(&wg)
	for i := 0; i < 10; i++ {
		go mapReader()
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

func mapReader() {
	valuesFound := []int{}
	for {
		val, found := sharedMap[rand.Intn(1000000)]
		if found {
			valuesFound = append(valuesFound, val)
		}
	}
}

// type bigStruct struct {
// 	data     []string
// 	children []bigStruct
// }

// var bigMap = make(map[string]bigStruct)
// var currentKey string

// func main() {
// 	for i := 0; i < 1000000; i++ {
// 		bigMap[fmt.Sprintf("key%d", i)] = bigStruct{}
// 	}

// 	wg := sync.WaitGroup{}
// 	wg.Add(1)
// 	go mapWriter(&wg)
// 	for i := 0; i < 10; i++ {
// 		go mapReader()
// 	}
// 	wg.Wait()
// }

// func mapWriter(wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for i := 0; i < 10000000; i++ {
// 		rInt := rand.Intn(1000000)
// 		rIntString := fmt.Sprintf("%d", rInt)
// 		currentKey := fmt.Sprintf("key%v", rIntString)
// 		val := bigStruct{[]string{rIntString, rIntString, rIntString, rIntString, rIntString, rIntString, rIntString, rIntString, rIntString, rIntString},
// 			[]bigStruct{
// 				{[]string{rIntString, rIntString, rIntString, rIntString, rIntString, rIntString, rIntString, rIntString, rIntString, rIntString}, nil},
// 				{[]string{rIntString, rIntString, rIntString, rIntString, rIntString, rIntString, rIntString, rIntString, rIntString, rIntString}, nil},
// 				{[]string{rIntString, rIntString, rIntString, rIntString, rIntString, rIntString, rIntString, rIntString, rIntString, rIntString}, nil},
// 			},
// 		}
// 		bigMap[currentKey] = val
// 	}
// }

// func mapReader() {
// 	for {
// 		val, found := bigMap[currentKey]
// 		if found {
// 			checkBigStruct(val)
// 		}
// 	}
// }

// func checkBigStruct(val bigStruct) {
// 	theString := val.data[0]

// 	data := val.data
// 	for _, v := range data {
// 		if v != theString {
// 			panic(fmt.Sprintf("Found %v, expected %v in %v", v, theString, val))
// 		}
// 	}
// 	children := val.children
// 	for _, child := range children {
// 		checkBigStruct(child)
// 	}
// }
