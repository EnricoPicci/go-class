# Cancel pattern

In this example we implement the cancel pattern.

The example starts with a main program that implements the forkJoin pattern. It launches 3 subtasks (i.e. 3 goroutines):

- slow-cooy: reads a file a copy its content into a strig - the read is slow since we use an ad hoc SlowReader that slows down read
- slow-file-processing: it reads from a file line by line and simulates a slow processing for each line read
- faulty-function: a function errors after a configurable period of time
  The 3 goroutines are passed a context.

Main waits for the partial results from the goroutines on a result channel. If any of the subtasks errors, the error is sent to an error channel.
If an error is received on the error channel, then the context is cancelled and the other subtasks terminate their processing gracefully.

## build

From the GO-CLASS project folder run the command
`go build -o ./bin/cancel-pattern ./src/concurrency-patterns/cancel-pattern`

## run

### run generating an error that triggers the cancel pattern

In order to trigger the cancel pattern it is necessary that `faultyFunction` fails. This occurs if we pass a time using the -failAfter flag.

- A very short time (e.g. <10) raise the error before both `slowlyReadFileToString` and `slowlyProcess` complete.
- A longer time (e.g. 1000) raise the error after `slowlyReadFileToString` completes but before `slowlyProcess` completes.
- An even longer time (e.g. 2000) raise the error after both `slowlyReadFileToString` and `slowlyProcess` complete.

From the GO-CLASS project folder run the command
`./bin/cancel-pattern -failAfter 100`

### run with no error

From the GO-CLASS project folder run the command
`./bin/cancel-pattern`
