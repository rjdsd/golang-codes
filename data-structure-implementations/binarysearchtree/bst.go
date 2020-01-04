package binarysearchtree

import (
	"fmt"
	"sync"
)

type Item struct {
	value int   // all types satisfy the empty interface, so we can store anything here.
	left *Item  // link to the next Item
	right *Item // link to the prev Item
}

type bstinterface interface {
	NewBST() *BinarySearchTree
    Size() int
	InsertItem(item *Item)
	DeleteItem(item *Item) 
	InorderTraversal()
	LowestCommonAnchestor(item1 *Item, item2 *Item) *Item
}

type BinarySearchTree struct {
	size int 
	root *Item // link to the next Item
	rwLock  sync.RWMutex
}

func NewBST() *BinarySearchTree {
	bst := new(BinarySearchTree)
	bst.size = 0
	bst.rwLock = *new(sync.RWMutex)
	return bst
}

func (bst *BinarySearchTree) Size() int{
	return bst.size
}

func (bst *BinarySearchTree) InsertItem(item *Item) {
	bst.rwLock.Lock()
	if bst.root == nil {
		bst.root = item
		defer bst.rwLock.Unlock()
		return
	} else {
			cur := bst.root
			for {
				if item.value <= cur.value {
					if cur.left == nil {
						cur.left = item
						defer bst.rwLock.Unlock()
						return
					} else {
						cur = cur.left
					}	
				} else {
						if cur.right == nil {
						cur.right = item
						defer bst.rwLock.Unlock()
						return
						} else {
						cur = cur.right
						}	

				}
			}
	}
}

func (bst *BinarySearchTree) InorderTraversal() {
	fmt.Println("===Inorder Traversal===")
	if bst.root == nil {
		return
	} else {
		doInorderTraversal(bst.root)
	}
	fmt.Println("end")
}

func  doInorderTraversal(cur *Item) {
	if cur == nil {
		return
	} else {
		doInorderTraversal(cur.left)
		fmt.Print(cur.value)
		fmt.Print("-->")
		doInorderTraversal(cur.right)
	}
}

func (bst *BinarySearchTree) LowestCommonAnchestor(item1 *Item, item2 *Item) *Item {
	if bst.root == nil {
		return nil
	} else {
		return findLCA(bst.root, item1, item2)
	}
}

func  findLCA(cur *Item, item1 *Item, item2 *Item) *Item {
	if cur == nil {
		return nil
	} else if cur.left == nil && cur.right == nil && cur.value != item1.value && cur.value != item2.value {
		return nil
	} else if cur.value < item1.value && cur.value < item2.value {
		return findLCA(cur.right, item1, item2)
	} else if cur.value > item1.value && cur.value > item2.value {
		return findLCA(cur.left, item1, item2)
	} else {
		return cur
	}
}
