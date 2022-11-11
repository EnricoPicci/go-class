# All votes go to the right candidate

To understand the starting point read the `readme.md` in the [voters](../voters/) and [voters](../voters-vote-synchronized/) folders.

We synchronize the call to the Vote method as well as the assignement to the interface type variable.

If we do this all data races are removed and the number of votes is exactly what we expetc: 1 Mil for each candidate.

## build

From the GO-CLASS project folder run the command
`go build -o ./bin/voters-fully-synchronized ./src/synchronization/data-race-with-interface/voters-fully-synchronized`

## build for data races detection

From the GO-CLASS project folder run the command
`go build -race -o ./bin/voters-fully-synchronized ./src/synchronization/data-race-with-interface/voters-fully-synchronized`

### run

From the GO-CLASS project folder run the command
`./bin/voters-fully-synchronized`
