package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// In this example we show the different behaviors of sync.Mutex and sync.RWMutex in a scenario with many readers and few writers.
// The idea is to simulate an app which has users. The users can receive "like"s from other users.
// Any user that enters the app sees the number of "like" of another users.
// Users entering the app and reading the number of "like"s are simulated by a many goroutines (i.e. many reads).
// The one user sending the "like" is simulated by just one goroutine (i.e. few writes).
// All the goroutines, both the readers and the writer, start at a random time within a specified time window

// Parameters used in the simulation

// timeWindow: the time window in which the goroutines are kicked off
var timeWindow = 1000

// returns the moment a goroutine is kicked off - it is a random moment within the time window
func goroutineKickOfTime() time.Duration {
	return time.Duration(rand.Intn(timeWindow)) * time.Millisecond
}

// delay: an arbitrary delay added to both the read and the write operations to make the differences between Mutes and RWMutexes more visible
// without such delay the operations would be so fast that the difference would not be perceived
var delay = time.Duration(5) * time.Millisecond

// number of readers and writers
var numberOfReaders = 1000
var numberOfWriters = 1

// First we define 2 struct describing a User, one protected by a Mutex and one protected by a RWMutex.
type user_M struct {
	name  string
	likes int
	// add a mutex to the user struct to protect the likes field
	mu sync.Mutex
}

func (u *user_M) like() {
	// we protect the update of the likes field with the Lock method
	u.mu.Lock()
	u.likes++
	u.mu.Unlock()
}
func (u *user_M) readLikes() int {
	// we protect the read of the likes field with the same mutex we use for the update
	// since this is just a read operation
	u.mu.Lock()
	// add an arbitrary delay make the differences while dealing with reads between Mutex and RWMutes more visible.
	time.Sleep(delay)
	ret := u.likes
	u.mu.Unlock()
	return ret
}

type user_RWM struct {
	name  string
	likes int
	// add a mutex to the user struct to protect the likes field
	mu sync.RWMutex
}

func (u *user_RWM) like() {
	// we protect the update of the likes field with the Lock method
	u.mu.Lock()
	u.likes++
	u.mu.Unlock()
}
func (u *user_RWM) readLikes() int {
	// we protect the read of the likes field with the same mutex we use for the update
	// since this is just a read operation we protect it with the RLock method
	u.mu.RLock()
	// add an arbitrary delay make the differences while dealing with reads between Mutex and RWMutes more visible.
	time.Sleep(delay)
	ret := u.likes
	u.mu.RUnlock()
	return ret
}

// it is convenient to have an Interface for all types of users so we can define a single runSimulation function
type likeReaderSender interface {
	like()
	readLikes() int
}

var lastLikesRead int // used just to avoid the compiler to cry because of the unused variable

func main() {
	fmt.Println("This example shows the different behaviors of sync.Mutex and sync.RWMutex in a scenario with many readers and few writers.")
	fmt.Println("sync.Mutex should introduce an higher delay (latency) than sync.RWMutex since there are many readers and few writers.")
	fmt.Println("")
	user_Mutex := user_M{name: "The Mutex user", likes: 0}
	user_RWMutex := user_RWM{name: "The RWMutex user", likes: 0}

	fmt.Println("Simulation with Mutex")
	fmt.Println("---------------------")
	runSimulation(&user_Mutex)

	fmt.Print("\n\n")
	fmt.Println("Simulation with RWMutex")
	fmt.Println("---------------------")
	runSimulation(&user_RWMutex)
}

func runSimulation(u likeReaderSender) {

	var writeTime time.Duration // moment when the write operation is performed

	start := time.Now()

	var wg_w sync.WaitGroup
	wg_w.Add(numberOfWriters)

	for i := 0; i < numberOfWriters; i++ {
		go func() {
			writeTime = goroutineKickOfTime()
			time.Sleep(writeTime)
			u.like()
			wg_w.Done()
		}()
	}

	var wg_r sync.WaitGroup
	wg_r.Add(numberOfReaders)
	for i := 0; i < numberOfReaders; i++ {
		go func() {
			time.Sleep(goroutineKickOfTime())
			likes := u.readLikes()
			lastLikesRead = likes
			wg_r.Done()
		}()
	}

	wg_w.Wait()
	wg_r.Wait()

	fmt.Printf("The last value of likes read is: %d\n", lastLikesRead)
	fmt.Printf("The time when the write occurs is: %s\n", writeTime)
	fmt.Printf("The time to execute the program is: %s\n", time.Since(start))
}
