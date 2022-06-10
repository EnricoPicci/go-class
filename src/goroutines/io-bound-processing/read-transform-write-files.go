package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func addLineNumber(path string) []string {

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

	numberOfLines := len(lines)
	lineNumberWidth := len(strconv.Itoa(numberOfLines))

	numberedLines := make([]string, numberOfLines)
	for i, line := range lines {
		numberedLines[i] = fmt.Sprintf("%*d %s", lineNumberWidth, i+1, line)
	}

	return numberedLines

}
