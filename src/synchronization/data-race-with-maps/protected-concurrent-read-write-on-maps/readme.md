# Protected concurrent read write on maps

Example of how concurrent reads and writes on a shared Map can be protected to avoid the runtime errors described [here](../../concurrent-read-write-on-maps/)

## Protection with sync.Mutex

The logic to protect the map with a Mutex is in the [mutex](./mutex/) folder.

## Protection with sync.RWMutex

The logic to protect the map with a RWMutex is in the [rwmutex](./rwmutex/) folder.

## Protection with sync.Map

The logic to protect the concurrent read/write operations using a sync.Map is in the [syncmap](./syncmap/) folder.

## The logic of the test

The 3 mechanisms (Mutex, RWMutex and sync.Map) are implemented in 3 different packages

- mutex
- rwmutex
- syncmap

For all 3 cases (Mutex, RWMutex and sync.Map) the logico of the test is the same and is implemented starting from the function ProtectedConcurrentReadWrite which is present in all the packages

- A map where the concurrent read/operations occur is created
- A writer goroutine is launched - this goroutine performs the write operations
- A number of reader goroutines are lauched
- Therefore we have 1 writer and n readers

The tests of the 3 mechanisms are launched sequentially by the `main` function in the [main.go](./main.go) file.

These are the parameters of the test

- rangeOfKeys: keys of the map are in the range from 0 to rangeOfKeys
- writes: number of writes to be performed by the writer
- reads: number of reads performed by each reader
- readGoroutines: number of readers (each reder is a goroutine)
- delayMilliseconds: each read and write is distanciated by this delay in milliseconds

These parameters have default values but can also be set from the command line.

### build

From the GO-CLASS project folder run the command

`go build -o ./bin/protected-concurrent-read-write-on-maps ./src/synchronization/data-race-with-maps/protected-concurrent-read-write-on-maps`

### run

From the GO-CLASS project folder run the command

`./bin/protected-concurrent-read-write-on-maps`
