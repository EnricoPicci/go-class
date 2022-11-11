# Voters vote but some votes gets to the wrong candidate

To understand the starting point read the `readme.md` in the [voters](../voters/) folder.

We synchronize the assignement to the interface type variable but this does not solve any problem.

The explanation is the following

1. The assignement to the `voter` variable (of interface type Voter) is protected by mutex which means that the 2 goroutines can not perform it concurrently.

2. The assignement to an interface type variable remains a 2 words operation and therefore it can happen that GoroutineDem sets the dynamic type of the interface type variable to be demCitizen but does not complete the operation while the GoroutineRep calls the vote method on the variable.
3. Now the variable holds a type demCitizen and runs the vote (for a repCitizen) with the demCitizen implementation
4. But now the GoroutineDem finishes the assignement operation and calls the vote method on the variable, so we have 2 concurrent execution of the same vote method implementation, one done by the GoroutineDem and one done by the GouroutineRep, and these 2 executions increment concurrently the same counter, i.e. the counter of the votes for Democratics
5. A concurrent write on the same share variable can corrupt data

## build

From the GO-CLASS project folder run the command

`go build -o ./bin/voters-assignement-synchronized ./src/synchronization/data-race-with-interface/voters-assignement-synchronized`

## build for data races detection

From the GO-CLASS project folder run the command

`go build -race -o ./bin/voters-assignement-synchronized ./src/synchronization/data-race-with-interface/voters-assignement-synchronized`

### run

From the GO-CLASS project folder run the command
`./bin/voters-assignement-synchronized`
