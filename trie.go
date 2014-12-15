package main

import (
	"bufio"
	"log"
	"os"
)

type Node struct {
	Children map[rune]*Node
	Words    map[string]struct{}
}

func (n *Node) Child(c rune) *Node {
	var _child *Node
	_child, ok := n.Children[c]
	if !ok {
		_child = NewNode()
		n.Children[c] = _child
	}
	return _child
}

func (n *Node) addWord(word string, curr int) {
	n.Words[word] = struct{}{}

	curr++
	if curr < len(word) {
		n.Child(rune(word[curr])).addWord(word, curr)
	}
}

func (n *Node) AddWord(word string) {
	substrs := Substrs(word, 2)
	var (
		_curr string
		_char rune
	)
	for i := 0; i < len(substrs); i++ {
		_curr = substrs[i]
		_char = rune(_curr[0])

		n.Child(_char).addWord(word, 0)
	}
}

func CreateTrie(filename string) *Node {
	top := NewNode()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var word string
	for scanner.Scan() {
		word = scanner.Text()
		// go top.AddWord(word)
		top.AddWord(word)
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return top
}

func NewNode() *Node {
	return &Node{
		Children: make(map[rune]*Node),
		Words:    make(map[string]struct{}),
	}
}
