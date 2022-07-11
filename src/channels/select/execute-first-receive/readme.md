# Use select to wait for the first receive that can be executed

In this example different data are sent on different channels after different random delays.
The select statement ensures that the first data sent to any of the channels is received unblocking the execution of the main function

## build

From the GO-CLASS project folder run the command
`go build -o ./bin/execute-first-receive ./src/channels/select/execute-first-receive`

### run

From the GO-CLASS project folder run the command
`./bin/execute-first-receive`
