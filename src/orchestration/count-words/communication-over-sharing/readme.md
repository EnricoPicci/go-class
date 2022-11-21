# Count words - Communication over sharing

These are examples of a programs that read the files in a directory and count the number of words in these files.

The logic of this example is based on the following steps:

- BuildDictionary function launches
  - One goroutine to read the file names contained in the directory
  - A pool of goroutines (readers) that read each file and calculate the number of words for this file
  - One goroutine that receives the partial results produced by the readers and accumulate them into a final result
- after launching these goroutines, BuildDictionary wait for the final resul on a channel
- the names of the files in the directory are read by a goroutine that sends each file name to a channel (fileChan) and then closes fileChan
- a pool of goroutines read from fileChan the file names one at a time; each goroutine builds a WordDictionary out of the file read and passes this dictionary to a channel (partialResultsChan); when all these goroutines have completed their work, partialResultsChan is closed
- another goroutine reads from partialResultsChan and accumulate each partial result into a WordDictionary the represents the final result; when all the partial results have been read and merged into the final result, then the final result is sent to a chan (finalResultsChan)
- BuildDictionary receives from finalResultsChan the final result and returns it

This example therefore uses a communication based design to complete the work. There are no shared variables used. On the contrary a "shared variables" based logic is implemented in the [sharingovercommunication](../sharing-over-communication/) package.

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
