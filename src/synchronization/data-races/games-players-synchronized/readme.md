# The Game Platform generates data races

Protecting the access to the shared resources (maps in this example) removes the data races

## build for daa races detection

From the GO-CLASS project folder run the command
`go build -race -o ./bin/games-players-synchronized ./src/synchronization/data-races/games-players-synchronized`

### run

From the GO-CLASS project folder run the command
`./bin/games-players-synchronized`
