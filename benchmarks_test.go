package main

import "testing"

const (
	listFilename      = "test/google-10000-english.txt"
	largeListFilename = "test/eowl-v1.1.2.txt"
)

func BenchmarkBuildTrie(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateTrie(listFilename)
	}
}

func benchmarkQuery(query string, b *testing.B) {
	trie := CreateTrie(largeListFilename)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trie.WordsContaining(query)
	}
}

func BenchmarkQueryEr(b *testing.B) { benchmarkQuery("er", b) }
