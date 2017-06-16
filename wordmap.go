package wordpatterns

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/mrap/stringutil"
)

type wordmap map[string][]string

type Wordmap struct {
	m           wordmap
	wordSubstrs map[string][]string
	opts        WordmapOptions
	ignoreChars map[rune]struct{}
}

type WordmapOptions struct {
	IgnoreCase   bool
	IgnoreChars  []rune
	MinSubstrLen int
}

func NewWordmap(opts *WordmapOptions) *Wordmap {
	if opts == nil {
		opts = &WordmapOptions{}
	}
	if opts.MinSubstrLen < 1 {
		opts.MinSubstrLen = 1
	}

	ignoreChars := make(map[rune]struct{})
	for _, c := range opts.IgnoreChars {
		if opts.IgnoreCase {
			c = unicode.ToLower(c)
		}
		ignoreChars[c] = struct{}{}
	}

	return &Wordmap{
		m:           make(wordmap),
		wordSubstrs: make(map[string][]string),
		opts:        *opts,
		ignoreChars: ignoreChars,
	}
}

func (wm *Wordmap) Has(word string) bool {
	_, exists := wm.wordSubstrs[word]
	return exists
}

func (wm *Wordmap) AddWord(word string) {
	substrs := stringutil.Substrs(wm.filteredSubstr(word), wm.opts.MinSubstrLen)
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
	return wm.m[wm.filteredSubstr(substr)]
}

func (wm Wordmap) SubstringCount() int {
	return len(wm.m)
}

func (wm Wordmap) filteredSubstr(substr string) string {
	if wm.opts.IgnoreCase {
		substr = strings.ToLower(substr)
	}
	return wm.removeIgnoredChars(substr)
}

func (wm Wordmap) removeIgnoredChars(str string) string {
	mapFunc := func(c rune) rune {
		if _, ignore := wm.ignoreChars[c]; ignore {
			return -1
		}
		return c
	}

	return strings.Map(mapFunc, str)
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
