package doublelinkedlist

import (
    "testing"
)

func TestList(t *testing.T) {
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
	dd.DeleteItem(&item2)
	dd.PrintList()
	item5 := Item{300.99,nil,nil}
	dd.DeleteItem(&item5)
	dd.PrintList()
}