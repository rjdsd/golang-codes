package trie

import (
	"fmt"
)

const alphabet_size = 26

type TrieNode struct {
	val string
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
	for _, c := range word {
		index := (c - 97) % 26;
		if cur.nodes[index] == nil {
			newNode := CreateNewTrie()
			newNode.val = string(c)
			cur.nodes[index] = newNode
			cur = newNode
		} else {
			cur = cur.nodes[index]
		}
	}
	cur.isEndOfWord = true
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

//  WIP
//  Return top 5 matching words starting with this pattern
func (trie *TrieNode) SearchWordsMatchingPattern (pattern string) []string {
	cur := trie
	result := make([]string, 1)
	for _, c := range pattern {
		index := (c - 97) % 26
		if(cur.nodes[index] == nil){
			return nil;
		} else {
			cur = cur.nodes[index]
		}
	}
	if cur.isEndOfWord {
		result[0] = pattern
	}
	cur.findFirstnChilds(pattern, result, 5)
	return result;
}

// WiP
func (t *TrieNode) findFirstnChilds(curStr string, res []string, n int) {
	if len(res) == n {
		return
	}
	curStr = curStr + string(t.val)
	if t.isEndOfWord {
		res = append(res, curStr )
	}
	if len(res) == n {
			return
	} else {
			for i := 0; i < alphabet_size; i++ {
				if t.nodes[i] != nil {
					t.nodes[i].findFirstnChilds(curStr,res,n)
				}
				if len(res) == n {
					return
				}
			}  
	}
}

func main(){
	t  := CreateNewTrie();
	t.InsertAWord("abc")
	t.InsertAWord("abcd")
	fmt.Println(t.SearchIfWordExists("abc"))
	fmt.Println(t.SearchIfWordExists("abcd"))
	fmt.Println(t.SearchWordsMatchingPattern("abb"))
}