package wordpatterns

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/mrap/goutil/slices"
)

type Ranks map[string]int

func RankByOccurences(trie *Node) Ranks {
	_ranks := make(Ranks)

	_wg := &sync.WaitGroup{}
	for _, c := range trie.Children {
		rankByOccurences(_ranks, c, _wg)
	}
	_wg.Wait()
	return _ranks
}

func rankByOccurences(_ranks Ranks, _n *Node, _wg *sync.WaitGroup) {
	_wg.Add(1)

	for _, c := range _n.Children {
		rankByOccurences(_ranks, c, _wg)
	}
	_ranks[_n.Str] = len(_n.Words)

	_wg.Done()
}

func (_wm Wordmap) Ranked() []string {
	sorted := make([]string, len(_wm.m))
	i := 0
	for w, _ := range _wm.m {
		sorted[i] = w
		i++
	}
	sort.Slice(sorted, func(i, j int) bool {
		return _wm.Compare(sorted[i], sorted[j]) < 0
	})
	slices.ReverseStrings(sorted)
	return sorted
}

// -1 Less than: less words or (if equal words) greater alphabetically
// 0  Equal: same word count and same letter count
// 1  More than: more words or (if equal words) less than alphabetically
func (_wm Wordmap) Compare(a, b string) int {
	aMatches := _wm.WordsContaining(a)
	bMatches := _wm.WordsContaining(b)
	if len(aMatches) > len(bMatches) {
		return 1
	} else if len(aMatches) < len(bMatches) {
		return -1
	} else {
		// Compare by character length
		if len(a) > len(b) {
			return 1
		} else if len(a) < len(b) {
			return -1
		} else {
			// Lastly compare alphabetically
			return strings.Compare(b, a)
		}
	}
}

func swap(words []string, a, b int) {
	words[a], words[b] = words[b], words[a]
}

func (_wm Wordmap) PrintRanks(limit int) {
	ranked := _wm.Ranked()
	PrintWordmapArr(_wm, ranked, limit)
}

func PrintWordmapArr(wm Wordmap, arr []string, limit int) {
	if limit <= 0 {
		limit = len(arr)
	}
	var str string
	for i := 0; i < limit; i++ {
		str = arr[i]
		fmt.Println(str, ":", len(wm.WordsContaining(str)))
	}
}
