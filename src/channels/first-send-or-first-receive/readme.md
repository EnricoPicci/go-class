# After a send/receive on an unbuffered channel, which goroutine executes the next line first?

The send/receive operation is the only atomic operation which governs channels. This means that the send/receive is a point of synchronization for the 2 goroutines (the one sending and the one receiving) in the sense that both have to wait for each other to complete the send and receive respectively.

After the send and receive operations complete there is no guarantee about which goroutine will execute the next line of code.

This can be seen running this example.

## build

From the GO-CLASS project folder run the command
`go build -o ./bin/first-send-or-first-receive ./src/channels/first-send-or-first-receive`

## run

From the GO-CLASS project folder run the command
`./bin/first-send-or-first-receive`
