# The Game Platform generates data races

Calling concurrently AddAndJoin and Ban generates data races and probably also a fatal error for concurrent writes on the same map

## build for daa races detection

From the GO-CLASS project folder run the command
`go build -race -o ./bin/games-players-data-race ./src/synchronization/data-races/games-players-data-race`

### run

From the GO-CLASS project folder run the command
`./bin/games-players-data-race`
