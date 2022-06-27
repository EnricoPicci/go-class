# Voters vote but some votes gets to the wrong candidate

To understand the starting point read the `readme.md` in the `voters` folder.

We synchronize the call to the Vote method but ths does not solve the problem. In this case the sum of the votes i 2 mlns, i.e. no vote is lost,
but the candidates do not get exatly 1 mln votes each, as expected, but a slightly different number.
This means that some democrats have voted the Rep candidate and viceversa, in a non predictable way though.

The explanation is the following

1. The assignement to the `voter` variable (of interface type Voter) is still not protected by mutex and so it may happen that the goroutine counting the republican votes assigns to the variable `voter` a concrete value of type RepCitizen just before the other goroutine, the one counting the democratic votes, invoke the `voter.Vote()` method.

2. If this happen the goroutine counting the democratic votes invoke the `voter.Vote()` method thinking that it is invoking the DemCitizen implementation of this method while in reality it ends up executing the RepCitizen implementation of this method.

3. The `voter.Vote()` code is protected by a mutex, which means that runs synchronously. This ensures that no parallel write to the vote counters is ever made. This explains why the sum of the votes of the 2 Candidates is exactly 2 mln, i.e. why no vote is lost.

4. The summary is that the 2 mln invocations of the `voter.Vote()` method run synchronously and the data of counters are not corrupted by parallel writes. At the same time there are cases where the goroutine counting the democratic votes executes the RepCitizen implementation of the `Vote()` method and viceversa. This explains why the 2 counters for the 2 candidates do not end up being 1 mln each, as expected.

Running with the data race detector on provides us with a warning of data races in the line of code

            voter = d

## build

From the GO-CLASS project folder run the command
`go build -o ./bin/voters-vote-synchronized ./src/synchronization/data-race-with-interface/voters-vote-synchronized`

## build for data races detection

From the GO-CLASS project folder run the command
`go build -race -o ./bin/voters-vote-synchronized ./src/synchronization/data-race-with-interface/voters-vote-synchronized`

### run

From the GO-CLASS project folder run the command
`./bin/voters-vote-synchronized`
