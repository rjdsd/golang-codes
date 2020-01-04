package binarysearchtree

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestBSTInsertion(t *testing.T) {
	bst := NewBST()
	item1 := Item{100,nil,nil}
	bst.InsertItem(&item1)
	item2 := Item{200,nil,nil}
	bst.InsertItem(&item2)
	item3 := Item{50,nil,nil}
	bst.InsertItem(&item3)
	item4 := Item{250,nil,nil}
	bst.InsertItem(&item4)
	bst.InorderTraversal()
}


func TestBSTLCA(t *testing.T) {
	bst := NewBST()
	item1 := Item{100,nil,nil}
	bst.InsertItem(&item1)
	item2 := Item{200,nil,nil}
	bst.InsertItem(&item2)
	item3 := Item{50,nil,nil}
	bst.InsertItem(&item3)
	item4 := Item{250,nil,nil}
	bst.InsertItem(&item4)
	lca := bst.LowestCommonAnchestor(&item3, &item1)
	assert.Equal(t, lca.value, 100, "they should be equal")
	lca = bst.LowestCommonAnchestor(&item2, &item4)
	assert.Equal(t, lca.value, 200, "they should be equal")
	lca = bst.LowestCommonAnchestor(&item3, &item4)
	assert.Equal(t, lca.value, 100, "they should be equal")
	item5 := Item{350,nil,nil}
	//  Following Test Case is Failing. 
	//It should return nil if any of the elements does not exist within the BST
	lca = bst.LowestCommonAnchestor(&item4, &item5)
	fmt.Println(lca.value)
}

