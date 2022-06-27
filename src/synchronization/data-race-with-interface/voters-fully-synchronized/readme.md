# Voters vote but some votes are missing

Calling concurrently AddAndJoin and Ban generates data races and probably also a fatal error for concurrent writes on the same map

## build for daa races detection

From the GO-CLASS project folder run the command
`go build -o ./bin/voters-synchronized-1 ./src/synchronization/data-race-with-interface/voters-synchronized-1`

### run

From the GO-CLASS project folder run the command
`./bin/voters-synchronized-1`
