package doublelinkedlist

import (
	"fmt"
	"sync"
)

type Item struct {
	value interface{} // All types satisfy the empty interface, so we can store anything here.
	next *Item // link to the next Item
	prev *Item // link to the prev Item
}

type ddlinkedlistinterface interface {
	NewDoubleLinkedList() *ddlinkedlist
    Size() int
	Append(item *Item) 
	InsertAtFront(item *Item)
	DeleteFromEnd() 
	PrintList()
	DeleteFromFront()
	AddAfterItem(target *Item, value *Item)
	DeleteAfterItem(target *Item, value *Item)
}

type ddlinkedlist struct {
	size int 
	head *Item // link to the next Item
	rwLock  sync.RWMutex
}

func NewDoubleLinkedList() *ddlinkedlist {
	dd := new(ddlinkedlist)
	dd.size = 0;
	return dd
}

func (dd *ddlinkedlist) Size() int{
	return dd.size
}

func (dd * ddlinkedlist) Append(item *Item){
	if dd.head == nil {
		dd.head = item
	} else {
		cur := dd.head
		for {
			if cur.next == nil {
				break
			}
			cur = cur.next
		}
		cur.next = item
		item.prev = cur
	}
	dd.size++
}

func (dd * ddlinkedlist) PrintList(){
	cur := dd.head
		for {
			if cur == nil {
				break
			}
			fmt.Println(cur.value)
			cur = cur.next
		}
}

func (dd * ddlinkedlist) AddAfterItem (target *Item, newNode *Item) {
	cur := dd.head
	for {
		if cur.value == target.value {
			break
		}
		cur = cur.next
	}
	newNode.next = cur.next
	cur.next.prev = newNode
	cur.next = newNode
	newNode.prev = cur
}
