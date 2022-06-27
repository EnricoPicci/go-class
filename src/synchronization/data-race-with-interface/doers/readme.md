# The Game Platform generates data races

Calling concurrently AddAndJoin and Ban generates data races and probably also a fatal error for concurrent writes on the same map

## build for daa races detection

From the GO-CLASS project folder run the command
`go build -o ./bin/data-race-with-interface ./src/synchronization/data-race-with-interface`

### run

From the GO-CLASS project folder run the command
`./bin/data-race-with-interface`
