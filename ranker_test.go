package main

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ranker", func() {
	const testFilename = "test/ranking_test_words.txt"
	var (
		_trie *Node
	)

	BeforeEach(func() {
		_trie = CreateTrie(testFilename)
	})

	Describe("Ranking most occurent substrings", func() {
		var _ranks Ranks
		BeforeEach(func() {
			_ranks = RankByOccurences(_trie)
		})

		It("should pair strings with number of occurences", func() {
			Ω(_ranks["th"]).Should(Equal(3))
			Ω(_ranks["he"]).Should(Equal(3))
			Ω(_ranks["the"]).Should(Equal(2))
			fmt.Println(_ranks)
		})
	})
})
