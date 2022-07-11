# What happens when 2 sends occur very closed to each other on 2 channels that are on a select statement

In this example data is sent on a channel ch_1 first and then, after a short interval, some more data is sent on ch_2.
Main in waiting on a select that has 2 cases: receive from ch_1 and receive from ch_2.
We want to see whether the case on ch_1 is always selected (since the first send is on ch_1) or whether there are cases when ch_2 is selected, since the time difference is minimal.

## build

From the GO-CLASS project folder run the command
`go build -o ./bin/very-closed-sends ./src/channels/select/very-closed-sends`

### run

From the GO-CLASS project folder run the command
`./bin/very-closed-sends`
