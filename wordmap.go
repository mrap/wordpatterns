package wordpatterns

import (
	"bufio"
	"log"
	"os"

	"github.com/mrap/stringutil"
)

type Wordmap struct {
	m wordmap
}

func NewWordmap() *Wordmap {
	return &Wordmap{
		m: make(wordmap),
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

type wordmap map[string][]string

func (wm *Wordmap) AddWord(word string) {
	substrs := stringutil.Substrs(word, 2)
	for _, s := range substrs {
		wm.m[s] = append(wm.m[s], word)
	}
}

func (wm Wordmap) WordsContaining(substr string) []string {
	return wm.m[substr]
}
