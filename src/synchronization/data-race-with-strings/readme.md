# Examples of data races that can occur when using variables of string type

A string is a 2 words data structure, therefore when a shared variable of type string is concurrently accessed (in write) by more than one goroutines data races can occur.

This is an example of data races with strings that lead to corrupted state and incoherent behaviour of the program.

This example has been inspired by [this blog](https://dave.cheney.net/2014/06/27/ice-cream-makers-and-data-races).

## build

From the GO-CLASS project folder run the command

`go build -o ./bin/data-race-with-strings ./src/synchronization/data-race-with-strings/`

## build for data races detection

From the GO-CLASS project folder run the command

`go build -race -o ./bin/data-race-with-strings ./src/synchronization/data-race-with-strings/`

### run

From the GO-CLASS project folder run the command

`./bin/data-race-with-strings`
