# Context aware scanning function

In this example we create a function that does some io work (it reads a file line by line).

The io package is not context-aware and so we create a custom function that is context aware and that is able to gracefully terminate if the context it has been passed is cancelled.

## build

From the GO-CLASS project folder run the command
`go build -o ./bin/context-aware-scanning-function ./src/context/custom-context-aware-functions/context-aware-scanning-function`

### run

From the GO-CLASS project folder run the command
`./bin/context-aware-scanning-function`
