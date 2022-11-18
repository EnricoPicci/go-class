package sharingovercommunication

import (
	"log"
	"sync"

	"github.com/EnricoPicci/go-class/src/orchestration/count-words/dictionary"
	"github.com/EnricoPicci/go-class/src/orchestration/count-words/dispatcher"
)

func CountUniqueWords(dirPath string, numOfReaders int, _log bool) int {
	return BuildDictionary(dirPath, numOfReaders, _log).NumberOfUniqueWords()
}

func CountTotalNumberOfWords(dirPath string, numOfReaders int, _log bool) int {
	return BuildDictionary(dirPath, numOfReaders, _log).TotalNumberOfWords()
}

func BuildDictionary(dirPath string, numOfReaders int, _log bool) *dictionary.WordDictionary {
	fileChan := make(chan string, numOfReaders)

	finalResult := dictionary.NewWordDictionary()

	go dispatcher.DispatchFiles(dirPath, fileChan, _log)

	var wg sync.WaitGroup
	wg.Add(numOfReaders)
	for i := 0; i < numOfReaders; i++ {
		go calculatePartialResults(finalResult, fileChan, &wg, _log)
	}
	wg.Wait()

	return finalResult
}

var mu sync.Mutex

func calculatePartialResults(finalResult *dictionary.WordDictionary, fileChan <-chan string, wg *sync.WaitGroup, _log bool) {
	defer wg.Done()

	for filePath := range fileChan {
		if _log {
			log.Printf(">>>>>>> file to calculate partial results %v received\n", filePath)
		}
		dict := dictionary.NewWordDictionary()
		dict.CountUniqueWords(filePath)
		mu.Lock()
		finalResult.Merge(dict)
		mu.Unlock()
	}
}
