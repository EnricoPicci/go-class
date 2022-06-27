package main

import (
	"fmt"
	"sync"
)

type Candidate struct {
	Name  string
	Votes int
}

var demCandidate = &Candidate{"Barack", 0}
var repCandidate = &Candidate{"GeorgeW", 0}

var strangeDemVotes = 0
var strangeRepVotes = 0

type Voter interface {
	Vote()
}

type DemCitizen struct {
	Id int
}

func (d DemCitizen) Vote() {
	if d.Id != 1 {
		fmt.Printf("I am the Democrat doing stuff as a %v\n", d.Id)
		strangeDemVotes++
	}
	demCandidate.Votes++
}

type RepCitizen struct {
	Id int
}

func (r RepCitizen) Vote() {
	if r.Id != 2 {
		fmt.Printf("I am the Republican doing stuff as a %v\n", r.Id)
		strangeRepVotes++
	}
	repCandidate.Votes++
}

func main() {

	var voter Voter

	var mu sync.Mutex
	// citizensVote := func(d Voter, wg *sync.WaitGroup) {
	// 	for i := 0; i < 1000000; i++ {
	// 		voter = d
	// 		mu.Lock()
	// 		voter.Vote()
	// 		mu.Unlock()
	// 	}
	// 	fmt.Println("DONE")
	// 	wg.Done()
	// }

	var wg sync.WaitGroup
	wg.Add(2)

	d_1 := DemCitizen{1}
	// go citizensVote(d_1, &wg)
	go func(d Voter, wg *sync.WaitGroup) {
		for i := 0; i < 1000000; i++ {
			voter = d
			mu.Lock()
			voter.Vote()
			mu.Unlock()
		}
		fmt.Println("DONE")
		wg.Done()
	}(d_1, &wg)

	d_2 := RepCitizen{2}
	// go citizensVote(d_2, &wg)
	go func(d Voter, wg *sync.WaitGroup) {
		for i := 0; i < 1000000; i++ {
			voter = d
			mu.Lock()
			voter.Vote()
			mu.Unlock()
		}
		fmt.Println("DONE")
		wg.Done()
	}(d_2, &wg)

	wg.Wait()
	fmt.Println(demCandidate.Votes)
	fmt.Println(repCandidate.Votes)
	fmt.Println(strangeDemVotes)
	fmt.Println(strangeRepVotes)
}
