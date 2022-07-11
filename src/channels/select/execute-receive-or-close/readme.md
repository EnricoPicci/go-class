# Use select to wait for the first receive that can be executed

In this example data is sent on a channel ch_1 after a random delays.
After a different random delay another channel, ch_2, is closed.
The select statement ensures that the first event (either the send of data over ch_1 or the closing of ch_2) is processed.

## build

From the GO-CLASS project folder run the command
`go build -o ./bin/execute-receive-or-close ./src/channels/select/execute-receive-or-close`

### run

From the GO-CLASS project folder run the command
`./bin/execute-receive-or-close`
