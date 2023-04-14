# Count words - sharing over communication

These are examples of a programs that read the files in a directory and count the number of words in these files.

The logic of this example is based on the following steps:

- The program starts with the [main function](../cmd/sharing-over-communication/) which launches the function [BuildDictionary](./count.go#BuildDictionary)
- BuildDictionary function launches
  - launches one goroutine to read the file names contained in the directory
  - then builds a dictionary that has to hold the final result
  - then launches a pool of goroutines (readers) that read each file and calculate the number of words for this file; the dictionary that has to hold the final result is passed to each of the readers; each reader merges its partial results into the finalResult dictionary
  - then waits for all the readers to complete; once all readers have completed, then the finalResult dictionary contains the result of the calculation

This example therefore uses a share-variable based design to complete the work. The finalResult dictionary is a variables shared among all readers goroutines and the main goroutine. On the contrary a "communication based" based logic is implemented in the [communicationoversharing](../communication-over-sharing/) package.

This approach shows performance which are worse than the one obtained with the "communication" based approach ([communicationoversharing](../communication-over-sharing/) package). This is due to the fact that the access to the shared dictionary variable must be synchronized with a Mutex, and synchronization creates latency. The difference in performance though can be perceived only when the size of the workers pool is low (e.g. just 1 worker). The bigger the size of the pool, the less relevant becomes the latency introduced by the synchronization.

## build

From the GO-CLASS project folder run the command

`go build -o ./bin/countWordsWithSharing ./src/orchestration/count-words/cmd/sharing-over-communication`

## build with race detector

From the GO-CLASS project folder run the command

`go build -race -o ./bin/countWordsWithSharing ./src/orchestration/count-words/cmd/sharing-over-communication`

### run the program

From the GO-CLASS project folder run the command

`./bin/countWordsWithSharing -dir canti-divina-commedia -readers 10 -verbose `

### run the benchmark

It is possible to run a benchmark to verify the different performance varying the number of goroutines in the pool

`go test ./src/orchestration/count-words/sharing-over-communication -bench=. -count 5 -run none`
