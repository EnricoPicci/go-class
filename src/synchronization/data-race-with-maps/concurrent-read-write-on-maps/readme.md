# Concurrent read write on maps

Example of how concurrent reads and writes on a shared Map can cause a runtime error since Map is not a concurrency-safe type and the Go compiler can raise a runtime error if concurrent reads and writes occur

### build

From the GO-CLASS project folder run the command

`go build -o ./bin/concurrent-read-write-on-maps ./src/synchronization/data-race-with-maps/concurrent-read-write-on-maps`

### run

From the GO-CLASS project folder run the command

`./bin/concurrent-read-write-on-maps`
