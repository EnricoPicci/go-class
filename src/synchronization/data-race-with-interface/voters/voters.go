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
	Id string
}

func (d DemCitizen) Vote() {
	// this is the DemCitizen implementation of the Vote method
	// if d.Id is not "Democrat" it means that a concrete value NOT of type DemCitizen is running this code
	// which is clearly something wrong
	if d.Id != "Democrat" {
		// fmt.Printf("I should be a Democrat but it seems I am a %v\n", d.Id)
		strangeDemVotes++
	}
	demCandidate.Votes++
}

type RepCitizen struct {
	Id string
}

func (r RepCitizen) Vote() {
	// this is the RepCitizen implementation of the Vote method
	// if d.Id is not "Republican" it means that a concrete value NOT of type RepCitizen is running this code
	// which is clearly something wrong
	if r.Id != "Republican" {
		// fmt.Printf("I should be a Republican but it seems I am a %v\n", r.Id)
		strangeRepVotes++
	}
	repCandidate.Votes++
}

func main() {

	var voter Voter

	citizensVote := func(d Voter, wg *sync.WaitGroup) {
		for i := 0; i < 1000000; i++ {
			voter = d
			voter.Vote()
		}
		fmt.Println("DONE")
		wg.Done()
	}

	var wg sync.WaitGroup
	wg.Add(2)

	d_1 := DemCitizen{"Democrat"}
	go citizensVote(d_1, &wg)

	d_2 := RepCitizen{"Republican"}
	go citizensVote(d_2, &wg)

	wg.Wait()
	fmt.Println(demCandidate.Votes)
	fmt.Println(repCandidate.Votes)
	fmt.Println(strangeDemVotes)
	fmt.Println(strangeRepVotes)
}
