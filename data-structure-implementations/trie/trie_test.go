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
    
	assert.True(t, trie.SearchIfWordExists("abc"))
    assert.True(t, trie.SearchIfWordExists("abcd"))
    assert.False(t, trie.SearchIfWordExists("hello"))

}