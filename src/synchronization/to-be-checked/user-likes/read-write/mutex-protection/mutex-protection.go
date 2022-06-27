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

type user struct {
	name  string
	likes int
	// add a mutex to the user struct to protect the likes field
	mu sync.Mutex
}

func (u *user) like() {
	// we protect the update of the likes field with a mutex
	u.mu.Lock()
	u.likes++
	u.mu.Unlock()
}
func (u *user) readLikes() int {
	// we protect the read of the likes field with the same mutex we use for the update
	// to protect the read is also important to avoid undesired behavior
	u.mu.Lock()
	// add an arbitrary delay to the read, while the write operation has no delay ans so can be considered it as immediate
	// Since we use a syncMutex, the total duration of the execution of all reads will be approximatively equal to the
	// number of goroutines launched for read times the delay added to the read
	time.Sleep(5 * time.Millisecond)
	ret := u.likes
	u.mu.Unlock()
	return ret
}

var users map[string]*user // map of users - pointers are used since each user is a unique value to be shared in the program

var lastLikesRead int // used just to avoid the compiler to cry because of the unused variable
var writeTime time.Duration

func main() {
	users = make(map[string]*user)
	users["The user liked"] = &user{name: "The user liked", likes: 0}

	// the same user is shared in the program among all the goroutines which are responsible to add likes to it
	theUserLiked := users["The user liked"]

	// users that send likes
	numberOfUsersSendingLikes_Goroutines := 1

	// users that read the likes field but do not send any like
	numberOfUsersReadingLikes_Goroutines := 1000

	start := time.Now()

	var wg_w sync.WaitGroup
	wg_w.Add(numberOfUsersSendingLikes_Goroutines)

	for i := 0; i < numberOfUsersSendingLikes_Goroutines; i++ {
		go func() {
			writeTime = time.Duration(rand.Intn(1000)) * time.Millisecond
			time.Sleep(writeTime)
			theUserLiked.like()
			wg_w.Done()
		}()
	}

	var wg_r sync.WaitGroup
	wg_r.Add(numberOfUsersReadingLikes_Goroutines)
	for i := 0; i < numberOfUsersReadingLikes_Goroutines; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			likes := theUserLiked.readLikes()
			lastLikesRead = likes
			wg_r.Done()
		}()
	}

	wg_w.Wait()
	wg_r.Wait()

	fmt.Printf("The value of likes should be equal to the number of users who have sent likes: %d\n",
		numberOfUsersSendingLikes_Goroutines)
	fmt.Printf("The real value of likes is: %d\n", theUserLiked.likes)
	fmt.Printf("The last value of likes read is: %d\n", lastLikesRead)
	fmt.Printf("The time when the write occurs is: %s\n", writeTime)
	fmt.Printf("The time to execute the program is: %s\n", time.Since(start))
}
