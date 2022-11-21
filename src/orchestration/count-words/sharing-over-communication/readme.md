# Count words - sharing over communication

These are examples of a programs that read the files in a directory and count the number of words in these files.

The logic of this example is based on the following steps:

BuildDictionary function

- launches one goroutine to read the file names contained in the directory
- then builds a dictionary that has to hold the final result
- then launches a pool of goroutines (readers) that read each file and calculate the number of words for this file; the dictionary that has to hold the final result is passed to each of the readers; each reader merges its partial results into the finalResult dictionary
- then waits for all the readers to complete; once all readers have completed, then the finalResult dictionary contains the result of the calculation

This example therefore uses a share-variable based design to complete the work. There finalResult dictionary is a variables shared among all readers goroutines nad the main goroutine. On the contrary a "communication based" based logic is implemented in the [communicationoversharing](../communication-over-sharing/) package.

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
