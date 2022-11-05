# Interfaces

## copy

In the [copy](./copy) folder there is an example of using the io.Copy function to copy from a Reader (the standard input) to a Writer (a file).

To run the example program use the following command

`go run ./src/interfaces/copy`

## mocking

In the [mocking](./mocking) folder there is an example of how to create a mock mechanism on the client decoupled from the api provider side.

Run the example against the real implementation (non mock) with the command

`go run ./src/interfaces/mocking/api-client`

Run the example against the mock implementation with the command
`go run ./src/interfaces/mocking/api-client -mock`

## reader-writer

In the [reader-writer](./reader-writer) folder there is an example of a program that copies from standard in to a file, either local or remote.

The copy logic is implemented using the io.Copy function since stdin and the target files implement respectively the io.Reader and io.Writer interface.
