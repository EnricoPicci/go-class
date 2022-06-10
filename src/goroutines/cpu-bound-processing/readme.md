# CPU bound processing example

## benchmark

From the GO-CLASS project folder run the command
`go test ./src/goroutines/cpu-bound-processing -bench=. -count 5 -run none`

## build

From the GO-CLASS project folder run the command
`go build -o ./bin/cpu-bound-processing ./src/goroutines/cpu-bound-processing`

## run

From the GO-CLASS project folder run the command
`./bin/cpu-bound-processing`
