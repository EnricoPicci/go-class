package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type fileWithLines struct {
	path  string
	lines []string
}

func readFilepathsFromDir(dirPath string) []string {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		err := fmt.Sprintf("read dir error: %v", err)
		fmt.Println(err)
		panic(err)
	}

	filePaths := make([]string, len(files))
	for i, file := range files {
		filePaths[i] = filepath.Join(dirPath, file.Name())
	}
	return filePaths
}

func calcFileRanges(files []string, numberOfRanges int) [][]string {
	numberOfFiles := len(files)
	fileRanges := make([][]string, numberOfRanges)

	for i := 0; i < numberOfRanges; i++ {
		start := i * numberOfFiles / numberOfRanges
		end := (i + 1) * numberOfFiles / numberOfRanges
		fileRanges[i] = files[start:end]
	}
	return fileRanges
}
func readFileLines(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		err := fmt.Sprintf("open file error: %v", err)
		fmt.Println(err)
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("scan file error: %v", err))
	}
	return lines
}

func addLineNumber(lines []string) []string {
	numberOfLines := len(lines)
	lineNumberWidth := len(strconv.Itoa(numberOfLines))

	numberedLines := make([]string, numberOfLines)
	for i, line := range lines {
		numberedLines[i] = fmt.Sprintf("%*d %s", lineNumberWidth, i+1, line)
	}

	return numberedLines
}

func addLineNumberToFiles(filePaths []string) []fileWithLines {
	fileWithLinesNumbered := make([]fileWithLines, len(filePaths))
	for i, filePath := range filePaths {
		lines := readFileLines(filePath)
		numberedLines := addLineNumber(lines)
		fileWithLinesNumbered[i] = fileWithLines{filePath, numberedLines}
	}
	return fileWithLinesNumbered
}

func write(file fileWithLines, dirPath string) {
	filePathSplit := strings.Split(file.path, "/")
	fileName := filePathSplit[len(filePathSplit)-1]
	outFilePath := filepath.Join(dirPath, fileName)
	f, err := os.Create(outFilePath)
	check(err)

	w := bufio.NewWriter(f)
	for _, line := range file.lines {
		_, err := w.WriteString(line + "\n")
		check(err)
	}

	w.Flush()
	//fmt.Printf("File %v written\n", file.path)

}

func addLineNumberToFilesAndWrite(filePaths []string, dirPath string, wg *sync.WaitGroup) {
	defer wg.Done()

	fileWithLinesNumbered := addLineNumberToFiles(filePaths)
	for _, fileWithLine := range fileWithLinesNumbered {
		write(fileWithLine, dirPath)
	}
}

func addLineNumbersToFilesInDir(dirPath string, outDirPath string, concurrent int) {
	filePaths := readFilepathsFromDir(dirPath)
	fileRanges := calcFileRanges(filePaths, concurrent)

	var wg sync.WaitGroup
	wg.Add(len(fileRanges))

	for _, fileRange := range fileRanges {
		go addLineNumberToFilesAndWrite(fileRange, outDirPath, &wg)
	}

	wg.Wait()
	//fmt.Println("Program terminating")
}
