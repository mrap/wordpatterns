package main

import (
	"sync"
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
