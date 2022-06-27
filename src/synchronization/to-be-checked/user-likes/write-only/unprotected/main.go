package main

import (
	"fmt"
	"sync"
)

type user struct {
	name  string
	likes int
}

func (u *user) like() {
	// unprotected update of the likes of the user
	u.likes++
}

var users map[string]*user // map of users - pointers are used since each user is a unique value to be shared in the program

func main() {
	users = make(map[string]*user)
	users["The user liked"] = &user{"The user liked", 0}

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
