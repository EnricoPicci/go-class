package main

import (
	"fmt"
	"sync"
)

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

var users map[string]*user // map of users - pointers are used since each user is a unique value to be shared in the program

func main() {
	users = make(map[string]*user)
	users["The user liked"] = &user{name: "The user liked", likes: 0}

	// the same user is shared in the program among all the goroutines which are responsible to add likes to it
	theUserLiked := users["The user liked"]

	numberOfUsersSendingLikes_Goroutines := 100000
	numberOfLikesPerUser := 1

	var wg sync.WaitGroup
	wg.Add(numberOfUsersSendingLikes_Goroutines)

	for i := 0; i < numberOfUsersSendingLikes_Goroutines; i++ {
		go func() {
			for j := 0; j < numberOfLikesPerUser; j++ {
				theUserLiked.like()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("The value of likes should be equal to the number of users who have sent likes and the number of likes per user: %d\n",
		numberOfUsersSendingLikes_Goroutines*numberOfLikesPerUser)
	fmt.Printf("The real value of likes is: %d\n", theUserLiked.likes)
}
