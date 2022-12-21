# Manage cancellation of multiple concurrent operations executing the cancel function of a context

This is an example where we implement an all-or-nothing kind of logic.

The program tries to execute a bunch of http get requests concurrently. If any of them fails, then all in-flight requests are cancelled.

The cancellation of the in-flight requests is acheived calling the cancel function returned by context.WithCancel API while creating the parent context, i.e. the context which is parent of all other contextes that are passed to the the http requests.

## build

From the GO-CLASS project folder run the command
`go build -o ./bin/with-cancel ./src/context/with-cancel`

### run

From the GO-CLASS project folder run the command
`./bin/with-cancel`
