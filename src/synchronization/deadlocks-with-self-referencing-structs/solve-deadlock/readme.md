# Scenario using 2 self referencing structs - how to solve the deadlock

The deadlock is eliminated simply reducing the scope of the mutexes to the bare minimum.

## build

From the GO-CLASS project folder run the command
`go build -o ./bin/self-referencing-structs-solve-deadlock ./src/synchronization/deadlocks-with-self-referencing-structs/solve-deadlock`

### run

From the GO-CLASS project folder run the command
`./bin/self-referencing-structs-solve-deadlock`
