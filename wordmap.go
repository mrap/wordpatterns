package wordpatterns

import (
	"bufio"
	"log"
	"os"

	"github.com/mrap/stringutil"
)

type wordmap map[string][]string

type Wordmap struct {
	m            wordmap
	MinSubstrLen int
	wordSubstrs  map[string][]string
}

func NewWordmap() *Wordmap {
	return &Wordmap{
		m:            make(wordmap),
		MinSubstrLen: 1,
		wordSubstrs:  make(map[string][]string),
	}
}

func (wm *Wordmap) Has(word string) bool {
	_, exists := wm.wordSubstrs[word]
	return exists
}

func (wm *Wordmap) AddWord(word string) {
	substrs := stringutil.Substrs(word, wm.MinSubstrLen)
	for _, s := range substrs {
		wm.m[s] = append(wm.m[s], word)
	}
	wm.wordSubstrs[word] = substrs
}

func (wm *Wordmap) RemoveWord(word string) {
	substrs := stringutil.Substrs(word, wm.MinSubstrLen)
	for _, s := range substrs {
		if words, exists := wm.m[s]; exists {
			wm.m[s] = removeStr(words, word)
		}
	}
	delete(wm.wordSubstrs, word)
}

func (wm Wordmap) WordsContaining(substr string) []string {
	return wm.m[substr]
}

func removeStr(arr []string, str string) []string {
	i := 0
	for _, s := range arr {
		if s != str {
			arr[i] = s
			i++
		}
	}
	return arr[:i]
}

func PopulateFromFile(wm *Wordmap, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wm.AddWord(scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
