# Manage cancellation of multiple concurrent operations executing the cancel function of a context

In this example we pass a context to one goroutine that has to perform an operation which expects a context, e.g. an http request.
Another concurrent goroutine is running another operation which ends up failing while the http request is still in flight,
we execute the cancel function on the parent context which cancels all children contexts passed down to the goroutines, which
causes the goroutine operations to terminate their own IO operations freeing up resources.

## build

From the GO-CLASS project folder run the command
`go build -o ./bin/with-timeout ./src/context/with-timeout`

### run

From the GO-CLASS project folder run the command
`./bin/with-timeout`
