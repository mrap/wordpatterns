package wordpatterns_test

import (
	. "github.com/mrap/wordpatterns"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Wordmap", func() {
	var wm *Wordmap

	BeforeEach(func() {
		wm = NewWordmap()
	})

	Describe("Adding a word", func() {
		const word = "abc"

		BeforeEach(func() {
			Expect(wm.Has(word)).To(BeFalse())
			wm.AddWord(word)
		})

		It("should have it before adding it", func() {
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
})
