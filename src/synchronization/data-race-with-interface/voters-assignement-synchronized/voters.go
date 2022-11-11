package main

import (
	"fmt"
	"sync"
)

const Democratic = "Democratic"
const Republican = "Republican"

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
	Party string
}

func (d DemCitizen) Vote() {
	// this is the DemCitizen implementation of the Vote method
	// if d.Party is not "Democrat" it means that a concrete value NOT of type DemCitizen is running this code
	// which is clearly something wrong
	if d.Party != Democratic {
		fmt.Printf("I should be a Democrat but it seems I am a %v\n", d.Party)
		strangeDemVotes++
	}
	demCandidate.Votes++
}

type RepCitizen struct {
	Party string
}

func (r RepCitizen) Vote() {
	// this is the RepCitizen implementation of the Vote method
	// if r.Party is not "Republican" it means that a concrete value NOT of type RepCitizen is running this code
	// which is clearly something wrong
	if r.Party != Republican {
		fmt.Printf("I should be a Republican but it seems I am a %v\n", r.Party)
		strangeRepVotes++
	}
	repCandidate.Votes++
}

func main() {

	var voter Voter

	var mu sync.Mutex
	citizensVote := func(d Voter, wg *sync.WaitGroup) {
		for i := 0; i < 1000000; i++ {
			mu.Lock()
			voter = d
			mu.Unlock()
			voter.Vote()
		}
		fmt.Println("DONE")
		wg.Done()
	}

	var wg sync.WaitGroup
	wg.Add(2)

	d := DemCitizen{Democratic}
	go citizensVote(d, &wg)

	r := RepCitizen{Republican}
	go citizensVote(r, &wg)

	wg.Wait()
	fmt.Println(demCandidate.Votes)
	fmt.Println(repCandidate.Votes)
	fmt.Println(strangeDemVotes)
	fmt.Println(strangeRepVotes)
}
