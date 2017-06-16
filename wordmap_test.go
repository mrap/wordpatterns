package wordpatterns_test

import (
	. "github.com/mrap/wordpatterns"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Wordmap", func() {
	var wm *Wordmap

	BeforeEach(func() {
		wm = NewWordmap(nil)
	})

	Describe("Adding a word", func() {
		const word = "abc"

		BeforeEach(func() {
			Expect(wm.Has(word)).To(BeFalse())
			wm.AddWord(word)
		})

		It("should have it after adding it", func() {
			Expect(wm.Has(word)).To(BeTrue())
		})

		It("should have word substrings", func() {
			Expect(wm.WordsContaining(word)).ToNot(BeEmpty())
		})
	})

	Describe("Removing a word", func() {
		const word = "abc"

		BeforeEach(func() {
			wm.AddWord(word)
			Expect(wm.Has(word)).To(BeTrue())
			wm.RemoveWord(word)
		})

		It("should not have it after removing it", func() {
			Expect(wm.Has(word)).To(BeFalse())
		})

		It("should not have word substrings", func() {
			Expect(wm.WordsContaining(word)).To(BeEmpty())
		})
	})

	Describe("Case sensitivity", func() {
		It("should be case sensitive by default", func() {
			wm := NewWordmap(nil)
			wm.AddWord("Ab")
			Expect(wm.WordsContaining("A")).To(HaveLen(1))
			Expect(wm.WordsContaining("a")).To(BeEmpty())
		})

		It("can be configured to be case insensitive", func() {
			wm := NewWordmap(&WordmapOptions{IgnoreCase: true})
			wm.AddWord("Ab")
			Expect(wm.WordsContaining("A")).To(HaveLen(1))
			Expect(wm.WordsContaining("a")).To(HaveLen(1))
		})
	})
})
