# Synchronization of a shared variable

Show how operations which operate on shared memory but are not correctly synchronized can cause problems

## Non sychronized shared counter

In this example a counter is incremented concurrently by many goroutines which run in parallel and therefore there is the high risk that, at the end of the processing, the counter will not hold the number expected just considering the number of times the increment operation has been executed

### build

From the GO-CLASS project folder run the command

`go build -o ./bin/non-synchronized-shared-counter ./src/synchronization/shared-values/shared-counter/non-synchronized-shared-counter`

### run

From the GO-CLASS project folder run the command

`./bin/non-synchronized-shared-counter`

### Non sychronized shared counter with just one processor

Is the same example as the one above with the only difference that we use a single processor so it is much less likely that we get an unexpected
number in the shared variable at the end of the processing. We say "much less likely" since the increment operation remains non atomic and therefore there is no guarantee that, for a number of iterations large enough, we may encounter the situation when the Scheduler decides to switch a goroutine when it is in the middle of updating the shared memory.

From the GO-CLASS project folder run the command

`./bin/non-synchronized-shared-counter -oneProc`

## Sychronized shared counter using atomic instructions

In this example a counter is incremented concurrently by many goroutines which run in parallel. Synchronization is guaranteed by the use of atomic instructions

### build

From the GO-CLASS project folder run the command

`go build -o ./bin/synchronized-shared-counter-atomic ./src/synchronization/shared-values/shared-counter/synchronized-shared-counter-atomic`

### run

From the GO-CLASS project folder run the command

`./bin/synchronized-shared-counter-atomic`

## Sychronized shared counter using mutexes

In this example a counter is incremented concurrently by many goroutines which run in parallel. Synchronization is guaranteed by the use of mutex

### build

From the GO-CLASS project folder run the command

`go build -o ./bin/synchronized-shared-counter-mutex ./src/synchronization/shared-values/shared-counter/synchronized-shared-counter-mutex`

### run

From the GO-CLASS project folder run the command

`./bin/synchronized-shared-counter-mutex`
