# Value of type interface with data races

In this example we show how a global shared variable of type interface can get corrupted if the assignement operation (which is a write operation) is done concurrently by more than one goroutine and is not protected by synchronization mechanisms such as Mutexes

## build

From the GO-CLASS project folder run the command
`go build -o ./bin/data-race-with-interface-doers ./src/synchronization/data-race-with-interface/interface-data-race-doers`

## build for data races

From the GO-CLASS project folder run the command
`go build -race -o ./bin/data-race-with-interface-doers ./src/synchronization/data-race-with-interface/interface-data-race-doers`

### run

From the GO-CLASS project folder run the command
`./bin/data-race-with-interface-doers`
