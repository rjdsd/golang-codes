package main

import (
	"fmt"
)

const alphabet_size = 26

type TrieNode struct {
	nodes [alphabet_size]*TrieNode
	isEndOfWord bool
}

type trie interface {
	CreateNewTrie() *TrieNode
	InsertAWord(word string)
	SearchIfWordExists(word string) bool
	SearchWordsMatchingPattern(pattern string) []string
}

func CreateNewTrie() *TrieNode{
	trie := new(TrieNode)
	return trie
}

func (trie *TrieNode) InsertAWord (word string){
	cur := trie;
	for i, c := range word {
		index := (c - 97) % 26;
		if cur.nodes[index] == nil {
			newNode := CreateNewTrie()
			cur.nodes[index] = newNode
			fmt.Println(cur)
			cur = newNode
		} 
		fmt.Println(i, " => ", c)
	}
	cur.isEndOfWord = true
	fmt.Println(cur)
}

func (trie *TrieNode) SearchIfWordExists(word string) bool {
	cur := trie
	for _, c := range word {
		index := (c - 97) % 26
		if cur.nodes[index] == nil {
			return false
		} else {
			cur = cur.nodes[index]
		}
	}
	if cur.isEndOfWord {
		return true
	}
	return false;
}

func (trie *TrieNode) SearchWordsMatchingPattern (pattern string) []string{
	return nil;

}

func main(){
	t  := CreateNewTrie();
	t.InsertAWord("abc")
	fmt.Println(t.SearchIfWordExists("hello"))
	fmt.Println(t.SearchIfWordExists("abc"))
}