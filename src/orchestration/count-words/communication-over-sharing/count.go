package communicationoversharing

import (
	"log"
	"sync"

	"github.com/EnricoPicci/go-class/src/orchestration/count-words/internal/dictionary"
	"github.com/EnricoPicci/go-class/src/orchestration/count-words/internal/dispatcher"
)

func CountUniqueWords(dirPath string, numOfReaders int, _log bool) int {
	return BuildDictionary(dirPath, numOfReaders, _log).NumberOfUniqueWords()
}

func CountTotalNumberOfWords(dirPath string, numOfReaders int, _log bool) int {
	return BuildDictionary(dirPath, numOfReaders, _log).TotalNumberOfWords()
}

func BuildDictionary(dirPath string, numOfReaders int, _log bool) *dictionary.WordDictionary {
	fileChan := make(chan string, numOfReaders)
	partialResultsChan := make(chan *dictionary.WordDictionary, numOfReaders)
	finalResultsChan := make(chan *dictionary.WordDictionary)

	go dispatcher.DispatchFiles(dirPath, fileChan, _log)

	go launchReaders(fileChan, partialResultsChan, numOfReaders, _log)

	go calculateFinalResult(partialResultsChan, finalResultsChan, _log)

	finalResult := <-finalResultsChan
	return finalResult
}

func launchReaders(fileChan <-chan string, partialResultsChan chan<- *dictionary.WordDictionary, numOfReaders int, _log bool) {
	var wg sync.WaitGroup
	wg.Add(numOfReaders)
	for i := 0; i < numOfReaders; i++ {
		go calculatePartialResults(fileChan, partialResultsChan, &wg, _log)
	}

	wg.Wait()
	if _log {
		log.Println(">>>>>>> partialResultsChan chan closed")
	}
	close(partialResultsChan)
}

func calculatePartialResults(fileChan <-chan string, partialResultsChan chan<- *dictionary.WordDictionary, wg *sync.WaitGroup, _log bool) {
	defer wg.Done()

	for filePath := range fileChan {
		if _log {
			log.Printf(">>>>>>> file to calculate partial results %v received\n", filePath)
		}
		dict := dictionary.NewWordDictionary()
		dict.CountUniqueWords(filePath)
		partialResultsChan <- dict
		if _log {
			log.Printf(">>>>>>> partial results for file %v sent\n", filePath)
		}
	}
}

func calculateFinalResult(partialResultsChan <-chan *dictionary.WordDictionary, finalResultsChan chan<- *dictionary.WordDictionary, _log bool) {
	finalResult := dictionary.NewWordDictionary()
	for dict := range partialResultsChan {
		if _log {
			log.Printf(">>>>>>> Partial results received with number of unique words %v\n", dict.NumberOfUniqueWords())
		}
		finalResult.Merge(dict)
	}

	finalResultsChan <- finalResult
	close(finalResultsChan)
}
