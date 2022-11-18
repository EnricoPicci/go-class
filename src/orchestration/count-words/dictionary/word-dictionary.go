package dictionary

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

type WordDictionary struct {
	dictionary map[string]int
	filesRead  []string
}

func NewWordDictionary() *WordDictionary {
	return &WordDictionary{dictionary: make(map[string]int)}
}
func (d *WordDictionary) Occurences(word string) int {
	count := d.dictionary[strings.ToLower(word)]
	return count
}

func (d *WordDictionary) CountUniqueWords(filePath string) {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		for _, word := range fields {
			word_ := strings.ToLower(word)
			word_ = removePunctuation(word_)
			if word_ != "" {
				d.dictionary[word_]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	d.filesRead = append(d.filesRead, filePath)
}

// var punctuationChars = ".,:;?!\""
var re = regexp.MustCompile(`[^\w\s]`)

func removePunctuation(word string) string {
	return re.ReplaceAllString(word, "")
}

func (d *WordDictionary) NumberOfUniqueWords() int {
	return len(d.dictionary)
}

func (d *WordDictionary) TotalNumberOfWords() int {
	totalNumberOfWords := 0
	for _, v := range d.dictionary {
		totalNumberOfWords = totalNumberOfWords + v
	}
	return totalNumberOfWords
}

// merges source into the receiver
func (d *WordDictionary) Merge(source *WordDictionary) {
	for k, v := range source.dictionary {
		wordCount, found := d.dictionary[k]
		if !found {
			d.dictionary[k] = v
		}
		d.dictionary[k] = wordCount + v
	}

	d.filesRead = append(d.filesRead, source.filesRead...)
}

func (d *WordDictionary) FilesRead() []string {
	return d.filesRead
}
func (d *WordDictionary) UniqueWords() []string {
	words := make([]string, len(d.dictionary))
	i := 0
	for k := range d.dictionary {
		words[i] = k
		i++
	}
	sort.Strings(words)
	return words
}
