package dictionary

import (
	"bufio"
	"fmt"
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

type Word struct {
	Value        string
	Occurrencies int
}

func (d *WordDictionary) UniqueWordsWithOccurrences() []Word {
	words := d.UniqueWords()
	resp := make([]Word, len(words))
	for i, v := range words {
		wordWithOccurrencies := Word{v, d.dictionary[v]}
		resp[i] = wordWithOccurrencies
	}
	return resp
}

func (d *WordDictionary) UniqueWordsSortedByOccurrences() []Word {
	wordsByOccurrencies := make([]string, len(d.dictionary))
	i := 0
	for k := range d.dictionary {
		wordsByOccurrencies[i] = k
		i++
	}
	sort.SliceStable(wordsByOccurrencies, func(i, j int) bool {
		return d.dictionary[wordsByOccurrencies[i]] > d.dictionary[wordsByOccurrencies[j]]
	})

	resp := make([]Word, len(wordsByOccurrencies))
	for i, v := range wordsByOccurrencies {
		wordWithOccurrencies := Word{v, d.dictionary[v]}
		resp[i] = wordWithOccurrencies
	}
	return resp
}

func (d *WordDictionary) PrintData(printWords bool, byOccurrencies bool, numWords int) {
	fmt.Printf("The number of files read is %v\n", len(d.FilesRead()))
	fmt.Printf("The total number of words is %v\n", d.TotalNumberOfWords())
	fmt.Printf("The number of unique words is %v\n", d.NumberOfUniqueWords())

	if printWords {
		fmt.Print("\n")
		var wordsWithOccurrencies []Word
		switch {
		case byOccurrencies:
			wordsWithOccurrencies = d.UniqueWordsSortedByOccurrences()
		default:
			wordsWithOccurrencies = d.UniqueWordsWithOccurrences()
		}
		for i := 0; i < numWords; i++ {
			v := wordsWithOccurrencies[i]
			if strings.TrimSpace(v.Value) == "" {
				fmt.Println(">>>>>>>> empty word")
			}
			fmt.Printf("%s (%v), ", v.Value, v.Occurrencies)
		}
		fmt.Print("\n")
	}
}
