# An http server creates the context starting from the request context

In this example we build an http server which serves requests arriving at a certain endpoint.
The http.Request returns a context and this context is used to create the context the server passes to the functions that process the request.
Doing so we have the advantage of being able to control the cancel signals coming from the http layer. For instance, if a client (e.g. a browser) disconnects, then the context fires a cancel signal and we can gracefully terminate the execution of the functions serving the request saving so CPU cycles.

## build the server

From the GO-CLASS project folder run the command
`go build -o ./bin/from-request-server ./src/context/from-request`

## run the server and test the client

### generate a timeout error

In order to generate a timeout error start the server running the following command

From the GO-CLASS project folder run the command
`./bin/from-request-server`

and then open a browser and request http://localhost:8080/process.

The reason this generates a timeout is that the detault timeout is 1 millisecond, too little to complete the http get request that the processing function executes. In this case we see on the server console that the error is "Request cancelled".

### execute the request

If we want a successful execution of the request we have to run the following command

From the GO-CLASS project folder run the command
`./bin/from-request-server -timeout 1000`

and then open a browser and request http://localhost:8080/process.

A timeout of 1.000 milliseconds should be enough to complete the http request that the processing function executes.

### generate a cancel error

In order to generate a cancel error start the server running the following command

From the GO-CLASS project folder run the command
`./bin/from-request-server -timeout 100000 -waitCancel true`

and then open a browser and request http://localhost:8080/process.

The request issued by the browser remains waiting for the cancellation signal that is fired when the browser tab is closed. In this case we see on the server console that the error is "Request cancelled".

If the timeout is shorter than the time the user takes to close the browser, then the error reported is that ot "timeout".

## help for the parameters that can be passed to the server

From the GO-CLASS project folder run the command
`./bin/from-request-server -h`
