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
	DeleteItem(target *Item)
}

type ddlinkedlist struct {
	size int 
	head *Item // link to the next Item
	rwLock  sync.RWMutex
}

func NewDoubleLinkedList() *ddlinkedlist {
	dd := new(ddlinkedlist)
	dd.rwLock = *new(sync.RWMutex)
	dd.size = 0;
	return dd
}

func (dd *ddlinkedlist) Size() int{
	return dd.size
}

func (dd * ddlinkedlist) Append(item *Item){
	dd.rwLock.Lock()
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
	dd.rwLock.Unlock()
}

func (dd * ddlinkedlist) PrintList(){
	fmt.Println("==")
	cur := dd.head
		for {
			if cur == nil {
				break
			}
			fmt.Print(cur.value)
			fmt.Print("-->")
			cur = cur.next
		}
		fmt.Println("nil")
}

func (dd * ddlinkedlist) AddAfterItem (target *Item, newNode *Item) {
	dd.rwLock.Lock()
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
	dd.rwLock.Unlock()
}

func (dd * ddlinkedlist) DeleteItem (target *Item) {
	dd.rwLock.Lock()
	cur := dd.head
	for {
		// It is important to check the nil first,
		// else when the target element
		// does not exist in the list, cur.value will throw exception
		if cur == nil || cur.value == target.value {
			break
		}
		cur = cur.next
	}
	if cur == dd.head {
		dd.head = dd.head.next
		dd.head.prev = nil
	} 
	if cur == nil {
		return
	} else {
		temp :=  cur.next
		cur.prev.next = temp
		if temp != nil {
			temp.prev = cur.prev
		}
	}
	dd.rwLock.Unlock()
}
