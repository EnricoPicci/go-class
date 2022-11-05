// go run ./src/interfaces/reader-writer/server/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

var logPtr *bool

var file *os.File

func createFileHandler(w http.ResponseWriter, r *http.Request) {
	if file != nil {
		errMsg := fmt.Sprintf("The file %v has not been closed by the previous client", file.Name())
		log.Println(errMsg)
		http.Error(w, errMsg, http.StatusNotAcceptable)
		return
	}
	type fileName struct {
		FileName string
	}
	f := &fileName{}
	err := json.NewDecoder(r.Body).Decode(f)
	if err != nil {
		log.Printf("Error %v while parsing body in the create request\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if *logPtr {
		log.Printf("File %v", *f)
	}

	fName := path.Join(".", "out", f.FileName)
	_file, err := os.Create(fName)
	if err != nil {
		errMsg := fmt.Sprintf("Error %v while creating file %v\n", err.Error(), f.FileName)
		log.Println(errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}
	file = _file

	filepath, err := filepath.Abs(file.Name())
	if err != nil {
		errMsg := fmt.Sprintf("Error %v while reading filePath for file %v\n", err.Error(), f.FileName)
		log.Println(errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	msg := fmt.Sprintf("file %v created", filepath)
	if *logPtr {
		log.Println(msg)
	}
	fmt.Fprint(w, msg)
}

func writeFileHandler(w http.ResponseWriter, r *http.Request) {
	if file == nil {
		errMsg := "Trying to write but there is no file open"
		log.Println(errMsg)
		http.Error(w, errMsg, http.StatusNotAcceptable)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error %v while reading body in the write request\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if *logPtr {
		log.Printf("Data written: %v", string(data))
	}

	n, err := file.Write(data)
	respData := map[string]any{"N": n, "Err": err}
	respJson, err := json.Marshal(respData)
	if err != nil {
		log.Printf("Error %v while marshalling write response %v\n", err.Error(), respData)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if *logPtr {
		log.Printf("RespJson sent after write '%v'\n", string(respJson))
	}
	_, err = w.Write(respJson)
	if err != nil {
		log.Printf("Error %v while sending %v after write\n", err.Error(), string(respJson))
		return
	}
}

func closeFileHandler(w http.ResponseWriter, r *http.Request) {
	if file == nil {
		errMsg := "Trying to close but there is no file open"
		log.Println(errMsg)
		http.Error(w, errMsg, http.StatusNotAcceptable)
		return
	}

	if *logPtr {
		log.Printf("Close file %v", file.Name())
	}

	err := file.Close()
	file = nil
	respData := map[string]any{"Err": err}
	respJson, err := json.Marshal(respData)
	if err != nil {
		log.Printf("Error %v while marshalling close response %v\n", err.Error(), respData)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if *logPtr {
		log.Printf("RespJson sent after close '%v'\n", string(respJson))
	}
	_, err = w.Write(respJson)
	if err != nil {
		log.Printf("Error %v while sending %v after close\n", err.Error(), string(respJson))
		return
	}
}

func main() {
	portPtr := flag.Int("port", 8081, "the port the server runs on")
	logPtr = flag.Bool("log", true, "log on the standard output")
	flag.Parse()

	http.HandleFunc("/create", createFileHandler)
	http.HandleFunc("/write", writeFileHandler)
	http.HandleFunc("/close", closeFileHandler)

	fmt.Printf("Starting server on port %v\n", *portPtr)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *portPtr), nil))
}
