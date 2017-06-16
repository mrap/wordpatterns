package wordpatterns

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/mrap/stringutil"
)

type wordmap map[string][]string

type Wordmap struct {
	m           wordmap
	wordSubstrs map[string][]string
	opts        WordmapOptions
}

type WordmapOptions struct {
	IgnoreCase   bool
	MinSubstrLen int
}

func NewWordmap(opts *WordmapOptions) *Wordmap {
	if opts == nil {
		opts = &WordmapOptions{}
	}
	if opts.MinSubstrLen < 1 {
		opts.MinSubstrLen = 1
	}

	return &Wordmap{
		m:           make(wordmap),
		wordSubstrs: make(map[string][]string),
		opts:        *opts,
	}
}

func (wm *Wordmap) Has(word string) bool {
	_, exists := wm.wordSubstrs[word]
	return exists
}

func (wm *Wordmap) AddWord(word string) {
	cased := word
	if wm.opts.IgnoreCase {
		cased = strings.ToLower(cased)
	}
	substrs := stringutil.Substrs(cased, wm.opts.MinSubstrLen)
	for _, s := range substrs {
		wm.m[s] = append(wm.m[s], word)
	}
	wm.wordSubstrs[word] = substrs
}

func (wm *Wordmap) RemoveWord(word string) {
	for _, s := range wm.wordSubstrs[word] {
		if words, exists := wm.m[s]; exists {
			wm.m[s] = removeStr(words, word)
		}
	}
	delete(wm.wordSubstrs, word)
}

func (wm Wordmap) WordsContaining(substr string) []string {
	if wm.opts.IgnoreCase {
		substr = strings.ToLower(substr)
	}
	return wm.m[substr]
}

func (wm Wordmap) SubstringCount() int {
	return len(wm.m)
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
