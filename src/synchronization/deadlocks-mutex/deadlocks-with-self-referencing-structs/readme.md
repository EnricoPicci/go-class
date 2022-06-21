# Dealock scenario using 2 self referencing structs

The simulation of a typical deadlock scenario with 2 user defined types which self reference each other and protect themselves via mutexes

## build

From the GO-CLASS project folder run the command
`go build -o ./bin/deadlocks-with-self-referencing-structs ./src/synchronization/deadlocks-mutex/deadlocks-with-self-referencing-structs`

### run

From the GO-CLASS project folder run the command
`./bin/deadlocks-with-self-referencing-structs`
