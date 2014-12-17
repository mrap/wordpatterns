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

func BenchmarkBuildWordmap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateWordmap(listFilename)
	}
}

func trieQuery(query string, b *testing.B) {
	trie := CreateTrie(largeListFilename)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trie.WordsContaining(query)
	}
}

func wordmapQuery(query string, b *testing.B) {
	wm := CreateWordmap(largeListFilename)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wm.WordsContaining(query)
	}
}

func BenchmarkTrieQueryEr(b *testing.B)    { trieQuery("er", b) }
func BenchmarkWordmapQueryEr(b *testing.B) { wordmapQuery("er", b) }
