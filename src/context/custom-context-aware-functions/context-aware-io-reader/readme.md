# Context aware io reader

In this example we create a Reader which is context aware. The reader is passed a context at construction time.
If the context is cancelled the reader stops reading.

This can be used, for instance, in cases when an external request comes in and requires to read a long file and copy it somewhere else. When the context is cancelled, the reader stops reading.

Example taken from

https://pace.dev/blog/2020/02/03/context-aware-ioreader-for-golang-by-mat-ryer.html

## build

From the GO-CLASS project folder run the command
`go build -o ./bin/context-aware-io-reader ./src/context/custom-context-aware-functions/context-aware-io-reader`

### run

From the GO-CLASS project folder run the command
`./bin/context-aware-io-reader`
