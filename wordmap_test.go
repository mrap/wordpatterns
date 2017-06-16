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
	})
})
