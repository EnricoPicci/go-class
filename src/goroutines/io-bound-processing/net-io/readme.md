# Network-IO bound processing example

Example of network IO bound processing.

The app sends a series of http requests to a web server.

The web server simulates some working by sleeping for some time for each request received.

<!-- The api of this transformation is the function

`addLineNumbersToFilesInDir` -->

## server

### build the server

From the GO-CLASS project folder run the command
`go build -o ./bin/net-io-bound-processing-server ./src/goroutines/io-bound-processing/net-io/server`

### run the server

From the GO-CLASS project folder run the command
`./bin/net-io-bound-processing-server`

## client

### build the client

From the GO-CLASS project folder run the command
`go build -o ./bin/net-io-bound-processing-client ./src/goroutines/io-bound-processing/net-io/client`

### run the client

From the GO-CLASS project folder run the command
`./bin/net-io-bound-processing-client`

## benchmark

Start the server with the command described above.

From the GO-CLASS project folder run the command
`go test ./src/goroutines/io-bound-processing/net-io/client -bench=. -count 5 -run none`
