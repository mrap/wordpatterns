package wordpatterns

import (
	"bufio"
	"log"
	"os"

	"github.com/mrap/stringutil"
)

type wordmap map[string][]string

type Wordmap struct {
	m           wordmap
	wordSubstrs map[string][]string
}

func NewWordmap() *Wordmap {
	return &Wordmap{
		m:           make(wordmap),
		wordSubstrs: make(map[string][]string),
	}
}

func CreateWordmap(filename string) *Wordmap {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wm := NewWordmap()
	for scanner.Scan() {
		wm.AddWord(scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return wm
}

func (wm *Wordmap) Has(word string) bool {
	_, exists := wm.wordSubstrs[word]
	return exists
}

func (wm *Wordmap) AddWord(word string) {
	substrs := stringutil.Substrs(word, 2)
	for _, s := range substrs {
		wm.m[s] = append(wm.m[s], word)
	}
	wm.wordSubstrs[word] = substrs
}

func (wm Wordmap) WordsContaining(substr string) []string {
	return wm.m[substr]
}
