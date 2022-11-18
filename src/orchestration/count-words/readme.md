# Count words

These are examples of a programs that read the files in a directory and count the number of words in these files.

## Communication over sharing

In the [communication-over-sharing](./communication-over-sharing/) folder there is an example of using a "communication over sharing" approach.

The logic is based on the following steps:

- BuildDictionary function launches
  - One goroutine to read the file names contained in the directory
  - A pool of goroutines (readers) that read each file and calculate the number of words for this file
  - One goroutine that receives the partial results produced by the readers and accumulate them into a final result
- after launching these goroutines, BuildDictionary wait for the final resul on a channel
- the names of the files in the directory are read by a goroutine that sends each file name to a channel (fileChan) and then closes fileChan
- a pool of goroutines read from fileChan the file names one at a time; each goroutine builds a WordDictionary out of the file read and passes this dictionary to a channel (partialResultsChan); when all these goroutines have completed their work, partialResultsChan is closed
- another goroutine reads from partialResultsChan and accumulate each partial result into a WordDictionary the represents the final result; when all the partial results have been read and merged into the final result, then the final result is sent to a chan (finalResultsChan)
- BuildDictionary receives from finalResultsChan the final result and returns it

This example therefore uses a communication based design to complete the work. There are no shared variables used. On the contrary a "shared variables" based logic is implemented in the [sharingovercommunication](./sharing-over-communication/) package.

## build

From the GO-CLASS project folder run the command

`go build -o ./bin/countWordsWithCommunication ./src/orchestration/count-words/cmd/communication-over-sharing`

## build with race detector

From the GO-CLASS project folder run the command

`go build -race -o ./bin/countWordsWithCommunication ./src/orchestration/count-words/cmd/communication-over-sharing`

### run the client

From the GO-CLASS project folder run the command

`./bin/countWordsWithCommunication -dir canti-divina-commedia -readers 10 -verbose `

### benchmark

Start the server with the command described above.

After the server has started, from the GO-CLASS project folder run the command

`go test ./src/orchestration/count-words/communication-over-sharing -bench=. -count 5 -run none`

## Sharing over Communication

In the [sharing-over-communication](./sharing-over-communication/) folder there is an example of using a "sharing over communication" approach.

The logic is based on the following steps:

BuildDictionary function

- launches one goroutine to read the file names contained in the directory
- then builds a dictionary that has to hold the final result
- then launches a pool of goroutines (readers) that read each file and calculate the number of words for this file; the dictionary that has to hold the final result is passed to each of the readers; each reader merges its partial results into the finalResult dictionary
- then waits for all the readers to complete; once all readers have completed, then the finalResult dictionary contains the result of the calculation

This example therefore uses a share-variable based design to complete the work. There finalResult dictionary is a variables shared among all readers goroutines nad the main goroutine. On the contrary a "communication based" based logic is implemented in the [communicationoversharing](./communication-over-sharing/) package.

## build

From the GO-CLASS project folder run the command

`go build -o ./bin/countWordsWithSharing ./src/orchestration/count-words/cmd/sharing-over-communication`

## build with race detector

From the GO-CLASS project folder run the command

`go build -race -o ./bin/countWordsWithSharing ./src/orchestration/count-words/cmd/sharing-over-communication`

### run the client

From the GO-CLASS project folder run the command

`./bin/countWordsWithSharing -dir canti-divina-commedia -readers 10 -verbose `

### benchmark

Start the server with the command described above.

After the server has started, from the GO-CLASS project folder run the command

`go test ./src/orchestration/count-words/sharing-over-communication -bench=. -count 5 -run none`
