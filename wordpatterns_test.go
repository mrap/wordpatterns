package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Wordpatterns", func() {
	const testFilename = "test_words.txt"
	var (
		_trie *Node
	)

	BeforeEach(func() {
		_trie = CreateTrie(testFilename)
	})

	Describe("Building a trie", func() {
		It("should correctly store 'the'", func() {
		})

		It("should build correctly", func() {
			// ->A
			a := _trie.Child('a')
			Ω(a.Words).ShouldNot(HaveKey("the"))
			Ω(a.Words).Should(HaveKey("that"))
			// ->A->T
			at := a.Child('t')
			Ω(at.Words).ShouldNot(HaveKey("the"))
			Ω(at.Words).Should(HaveKey("that"))

			// ->H
			h := _trie.Child('h')
			Ω(h.Words).Should(HaveKey("the"))
			Ω(h.Words).Should(HaveKey("that"))
			// ->H->A
			ha := h.Child('a')
			Ω(ha.Words).ShouldNot(HaveKey("the"))
			Ω(ha.Words).Should(HaveKey("that"))
			// ->H->A->T
			hat := ha.Child('t')
			Ω(hat.Words).ShouldNot(HaveKey("the"))
			Ω(hat.Words).Should(HaveKey("that"))
			// ->H->E
			he := h.Child('e')
			Ω(he.Words).Should(HaveKey("the"))
			Ω(he.Words).ShouldNot(HaveKey("that"))

			// ->T
			t := _trie.Child('t')
			Ω(t.Words).Should(HaveKey("the"))
			Ω(t.Words).Should(HaveKey("that"))
			// ->T->H
			th := t.Child('h')
			Ω(th.Words).Should(HaveKey("the"))
			Ω(th.Words).Should(HaveKey("that"))
			// ->T->H->E
			the := th.Child('e')
			Ω(the.Words).Should(HaveKey("the"))
			Ω(the.Words).ShouldNot(HaveKey("that"))
		})
	})
})
