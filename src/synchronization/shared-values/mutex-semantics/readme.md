# Pointer semantics for mutexes

Show how mutexes work if using pointer semantics

## Sychronized shared counter using pointer semantics

In this example a counter (value of type Counter) is incremented concurrently by many goroutines which run in parallel but the increment operation is protected by a Mutex. Pointer semantics is used with type Counter and therefore no data corruption occurs since the counter value is effectively protected by one single share value of type Mutex.

### build

From the GO-CLASS project folder run the command

`go build -o ./bin/mutex-pointer-semantics ./src/synchronization/shared-values/mutex-semantics/pointer-semantics`

### run

From the GO-CLASS project folder run the command

`./bin/mutex-pointer-semantics`

## Sychronized shared counter using VALUE semantics

In this example a counter (value of type Counter) is incremented concurrently by many goroutines which run in parallel, the increment operation is protected by a Mutex but VALUE semantics is used. This means that there are copies of the mutex value used in the various goroutines and therefore the Mutex can not protect effectively the counter from data corruption.

### build

From the GO-CLASS project folder run the command

`go build -o ./bin/mutex-value-semantics ./src/synchronization/shared-values/mutex-semantics/value-semantics`

### run

From the GO-CLASS project folder run the command

`./bin/mutex-value-semantics`
