package trie

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestStack(t *testing.T) {
    // assert equality
    assert.Equal(t, 123, 123, "they should be equal")


    trie  := CreateNewTrie();
	trie.InsertAWord("abc")
    trie.InsertAWord("abcd")
    trie.InsertAWord("abcdd")
    
	assert.True(t, trie.SearchIfWordExists("abc"))
    assert.True(t, trie.SearchIfWordExists("abcd"))
    assert.False(t, trie.SearchIfWordExists("helloworld"))
    
    trie.InsertAWord("helloworld")
    assert.True(t, trie.SearchIfWordExists("helloworld"))
    
    assert.Equal(t, 3, len(trie.SearchWordsMatchingPattern("abc")), "trie.SearchWordsMatchingPattern")
    assert.Equal(t, 1, len(trie.SearchWordsMatchingPattern("hello")), "trie.SearchWordsMatchingPattern")
    assert.Equal(t, 0, len(trie.SearchWordsMatchingPattern("world")), "trie.SearchWordsMatchingPattern")
}