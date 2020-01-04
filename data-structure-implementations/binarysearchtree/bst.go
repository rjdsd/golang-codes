package binarysearchtree

import (
	"fmt"
	"sync"
)

type Item struct {
	value int // All types satisfy the empty interface, so we can store anything here.
	left *Item // link to the next Item
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
	return bst
}

func (bst *BinarySearchTree) Size() int{
	return bst.size
}


func (bst *BinarySearchTree) InsertItem(item *Item) {
	if bst.root == nil {
		bst.root = item
		return
	} else {
		cur := bst.root
		for {
			if cur == nil {
				cur = item
				return
			} else {
				if item.value <= cur.value {
					cur = cur.left
				} else {
					cur = cur.right
				}
			}
		}
	}
}


func (bst *BinarySearchTree) InorderTraversal() {
	if bst.root == nil {
		return
	} else {
		doInorderTraversal(bst.root)
	}
}

func  doInorderTraversal(cur *Item) {
	if cur == nil {
		return
	} else {
		doInorderTraversal(cur.left)
		fmt.Println(cur)
		fmt.Println("---")
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
	} else if cur.value < item1.value && cur.value < item2.value {
		return findLCA(cur.right, item1, item2)
	} else if cur.value > item1.value && cur.value > item2.value {
		return findLCA(cur.left, item1, item2)
	} else {
		return cur
	}
}
