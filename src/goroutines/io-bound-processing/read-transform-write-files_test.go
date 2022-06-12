package main

import (
	"os"
	"strconv"
	"testing"

	"github.com/EnricoPicci/go-class.git/src/testhelpers"
)

func TestReadFilepathsFromDir(t *testing.T) {
	var path = testhelpers.FilePath("/canti-divina-commedia")

	files := readFilepathsFromDir(path)

	// test the number of files
	expectedNumberOfFiles := 100
	numberOfFiles := len(files)
	if numberOfFiles != expectedNumberOfFiles {
		t.Errorf("expected %d files, got %d", expectedNumberOfFiles, numberOfFiles)
	}

}

func TestCalcFileRanges(t *testing.T) {
	var path = testhelpers.FilePath("/canti-divina-commedia")
	filePaths := readFilepathsFromDir(path)

	var concurrent = 2
	fileRanges := calcFileRanges(filePaths, concurrent)

	expectedRanges := [][]string{
		{"01 - Inferno - CANTO PRIMO.txt", "02 - Inferno - CANTO SECONDO.txt"},
		{"01 - Inferno - CANTO PRIMO.txt", "02 - Inferno - CANTO SECONDO.txt"},
	}

	// test the number of ranges
	expectedNumberOfRanges := len(expectedRanges)
	numberOfRanges := len(fileRanges)
	if numberOfRanges != expectedNumberOfRanges {
		t.Errorf("The result is %v instead of %v", numberOfRanges, expectedNumberOfRanges)
	}

	// test the length of the ranges
	expectedRangeLength := 50
	for i, r := range fileRanges {
		if len(r) != 50 {
			t.Errorf("The range %v contains %v number of file instead of %v", i, len(r), expectedRangeLength)
		}
	}

	// create a set of file names contained in all the ranges (eliminating the duplicates) and check
	// that the number of files in the set is the same as the number of files read from the directory
	fileSet := make(map[string]struct{})
	for _, r := range fileRanges {
		for _, f := range r {
			fileSet[f] = struct{}{}
		}
	}
	numberOfUniqueFilesInRanges := len(fileSet)
	numberOfFilesReadFromFolder := len(readFilepathsFromDir(path))
	if numberOfUniqueFilesInRanges != numberOfFilesReadFromFolder {
		t.Errorf("The number of unique files %v is not the same of the number of files read from folde which is %v",
			numberOfUniqueFilesInRanges, numberOfFilesReadFromFolder)
	}

}

func TestAddLineNumber(t *testing.T) {
	var path = testhelpers.FilePath("/canti-divina-commedia/01 - Inferno - CANTO PRIMO.txt")

	lines := readFileLines(path)
	numberedLines := addLineNumber(lines)

	// test the number of lines
	expectedNumberOfLines := 137
	numberOfLines := len(numberedLines)
	if numberOfLines != expectedNumberOfLines {
		t.Errorf("The result is %v instead of %v", numberOfLines, expectedNumberOfLines)
	}

	// test the first number
	expectedFirstNumber := '1'
	firstLine := numberedLines[0]
	var firstNumber rune
	for _, rune := range firstLine {
		if rune != ' ' {
			firstNumber = rune
			break
		}
	}
	if firstNumber != expectedFirstNumber {
		t.Errorf("The result is %v instead of %v", string(firstNumber), string(expectedFirstNumber))
	}

	// test the last number
	expectedLastNumber := "137"
	lastLine := numberedLines[len(numberedLines)-1]
	lastNumber := lastLine[0:len(expectedLastNumber)]
	if lastNumber != expectedLastNumber {
		t.Errorf("The result is %v instead of %v", string(firstNumber), string(expectedLastNumber))
	}
}

func TestAddLineNumberToFiles(t *testing.T) {
	var path = testhelpers.FilePath("/canti-divina-commedia")
	filePaths := readFilepathsFromDir(path)

	var concurrent = 2
	fileRanges := calcFileRanges(filePaths, concurrent)

	numberedFiles := addLineNumberToFiles(fileRanges[0])

	// test the number of numberedFiles
	expectedNumberOfFiles := 50 // half of the files of the Divina Commedia
	if len(numberedFiles) != expectedNumberOfFiles {
		t.Errorf("The result is %v instead of %v", numberedFiles, expectedNumberOfFiles)
	}

	// test the last line of the first file starts with the right number
	firstFile := numberedFiles[0]
	numberOfLinesOfFirstFile := len(numberedFiles[0].lines)
	lastLine := firstFile.lines[numberOfLinesOfFirstFile-1]

	lastLineNumber := strconv.Itoa(numberOfLinesOfFirstFile)

	expectedLineNumber := lastLine[0:len(lastLineNumber)]
	if lastLineNumber != expectedLineNumber {
		t.Errorf("The result is %v instead of %v", lastLineNumber, expectedLineNumber)
	}
}

func TestWrite(t *testing.T) {
	var path = testhelpers.FilePath("canti-divina-commedia/01 - Inferno - CANTO PRIMO.txt")
	lines := readFileLines(path)

	// tmpDir := t.TempDir()
	// outFilePath := filepath.Join(tmpDir, "out.txt")
	outDir := testhelpers.FilePath("tmp/")
	outFileName := "01 - Inferno - CANTO PRIMO.txt"
	outFilePath := outDir + outFileName

	f := fileWithLines{path, lines}

	err := os.Remove(outFilePath)
	if err != nil {
		t.Logf("Error removing file: %v - maybe the file was just not present", err)
	}

	write(f, outDir)

	// test the number of lines written
	outFileLines := readFileLines(outFilePath)
	expectedNumberOfLines := len(lines)
	if len(outFileLines) != expectedNumberOfLines {
		t.Errorf("The lines written are %v instead of %v", len(outFileLines), expectedNumberOfLines)
	}
}

func TestAddLineNumbersToFilesInDir(t *testing.T) {
	var path = testhelpers.FilePath("/canti-divina-commedia")
	var outDirPath = t.TempDir()
	var concurrent = 2

	addLineNumbersToFilesInDir(path, outDirPath, concurrent)

	// test that the number of files in the output directory is the same of the number of files in the input directory
	files := readFilepathsFromDir(path)
	outFiles := readFilepathsFromDir(outDirPath)
	if len(files) != len(outFiles) {
		t.Errorf("The number of files in the output directory is %v instead of %v", len(outFiles), len(files))
	}

}
