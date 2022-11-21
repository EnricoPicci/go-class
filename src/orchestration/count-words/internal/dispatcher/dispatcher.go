package dispatcher

import (
	"log"
	"os"
	"path/filepath"
)

func DispatchFiles(dirPath string, fileChan chan<- string, _log bool) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if _log {
			log.Printf(">>>>>>> Trying to send file %v\n", file.Name())
		}
		fPath := filepath.Join(dirPath, file.Name())
		fileChan <- fPath
		if _log {
			log.Printf(">>>>>>> file %v sent to file chan\n", file.Name())
		}
	}

	if _log {
		log.Println(">>>>>>> File chan closed")
	}
	close(fileChan)
}
