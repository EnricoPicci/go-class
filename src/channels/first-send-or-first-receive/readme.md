# After a send/receive on an unbuffered channel, which goroutine executes the next line first?

The send/receive operation is the only atomic operation which governs channels.
The receive completes nanoseconds before the send but this does not guarantee which is the goroutine that will execute the next line of code.

## build

From the GO-CLASS project folder run the command
`go build -o ./bin/first-send-or-first-receive ./src/channels/first-send-or-first-receive`

### run

From the GO-CLASS project folder run the command
`./bin/first-send-or-first-receive`
