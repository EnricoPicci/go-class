# Voters vote but some votes are missing

2 candidates for the election are going to receive 1 mln votes each.
Each candidate receives its votes through a separate subroutine which loops 1 mln times and, for each loop, casts a vote to its candidate.
The loop logic looks like this

    	for i := 0; i < 1000000; i++ {
    		voter = d // d is a value of type DemCitizen or RepCitizen while voter is a variable of interface type Voter
    		voter.Vote()
    	}

At the end of the execution of the 2 goroutines none of the Candidates has 1 mln votes. Both have a slightly lower unpredictable number.

The reason is in the logic we have implemented. Let's see step by step

0. The voter variable is shared among the 2 goroutines. Both goroutines access the same variable.

1. In the voting loop we assign a value to an interface type variable in the line
   voter = d // d is a value of type DemCitizen or RepCitizen while voter is a variable of interface type Voter

2. With an interface type variable, the assignement operation is a 2 words operation, since an interface is a 2 words data structure (one word holds the pointer to the dynamic type and one word holds the pointer to the concrete value). Therefore it is not an atomic instruction.

3. Since the voter variable is shared among the 2 goroutines, it can happen that one goroutine reads the shared value while the other is in the middle of the assignement operation, which means for instance that

- Goroutine G1 (the one that casts the democratic votes) assigns a concrete value of type DemCitizen the `voter` variable
- Goroutine G2 (the one that casts the republican votes) starts the assignement of a concrete value of type RepCitizen the `voter` variable but sets just the type (RepCitizen) when G1 resume execution
- When G1 resumes execution it thinks that it is dealing with a DemCitizen but the variable `voter` (of type interface `Voter`) points to the type RepCitizen, i.e. the last type set by G2.
- At this point we have both G1 and G2 reach the call to the `voter.Vote()` method and both run in parallel the same implementation, i.e. the RepCitizen implmenetation of the `Vote()` method
- The `Vote()` method updates a piece of shared memory, in this case with the instruction `repCandidate.Votes++`
- As we know, if 2 or more goroutines run in parallel code that writes on shared memory there is no guarantee of the result (see the examples in the shared-value folder to see how parallel writes to share memory generate unpredictable results) and this explains why the number of the votes of each candidate is not 1 Mln

4.  The fact that sometimes a concrete value of type RepCitizen runs the code implemented for DemCitizen can be proved by the following code

            func (d DemCitizen) Vote() {
                // this is the DemCitizen implementation of the Vote method
                // if d.Id is not "Democrat" it means that a concrete value NOT of type DemCitizen is running this code
                // which is clearly something wrong
                if d.Id != "Democrat" {
                    fmt.Printf("I should be a Democrat but it seems I am a %v\n", d.Id)
                    strangeDemVotes++
                }
                demCandidate.Votes++
            }

If we run this simulation we see that there are cases when a concrete value NOT of type DemCitizen executes the DemCitizen implementation of `Vote` method. Same happens with RepCitizen.

In fact, at the end of the execution, the values of the variables `strangeDemVotes` and `strangeRepVotes` are greater than 0.

If we run the simulation with the data race detector we find that the data race detector identifies 3 points where data races occur, in the following lines of code

        - demCandidate.Votes++
        - repCandidate.Votes++
        - strangeDemVotes++
        - strangeRepVotes++
        - voter = d

Running the race detector warns us that there is something wrong here.

## build

From the GO-CLASS project folder run the command

`go build -o ./bin/voters ./src/synchronization/data-race-with-interface/voters`

## build for data races detection

From the GO-CLASS project folder run the command

`go build -race -o ./bin/voters ./src/synchronization/data-race-with-interface/voters`

### run

From the GO-CLASS project folder run the command

`./bin/voters`
