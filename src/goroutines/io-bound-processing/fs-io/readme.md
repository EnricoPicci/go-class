# FS-IO bound processing example

Example of file system IO bound processing.

The app reads files from a directory, add a line at the beginning of each line of each file and then writes the files with the lines numbered to an output directory.

The api of this transformation is the function

`addLineNumbersToFilesInDir`

## benchmark

From the GO-CLASS project folder run the command
`go test ./src/goroutines/io-bound-processing/fs-io -bench=. -count 5 -run none`

## test

From the GO-CLASS project folder run the command
`go test ./src/goroutines/io-bound-processing/fs-io`

## build

From the GO-CLASS project folder run the command
`go build -o ./bin/fs-io-bound-processing ./src/goroutines/io-bound-processing/fs-io`

## run

From the GO-CLASS project folder run the command
`./bin/fs-io-bound-processing`
