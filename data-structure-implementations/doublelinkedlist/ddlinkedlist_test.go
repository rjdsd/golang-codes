package doublelinkedlist

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
    // assert equality
	assert.Equal(t, 123, 123, "they should be equal")
	dd := NewDoubleLinkedList()
	item1 := Item{100,nil,nil}
	dd.Append(&item1)
	item2 := Item{400.75,nil,nil}
	dd.Append(&item2)
	item3 := Item{200,nil,nil}
	dd.AddAfterItem(&item1, &item3)
	item4 := Item{250,nil,nil}
	dd.AddAfterItem(&item3, &item4)
	dd.PrintList()
}