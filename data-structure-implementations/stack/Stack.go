/*
@author Rajdeep Sardar
This is an implementation of a Stack in Golang. This is non-BLocking.
*/

package stack
import "errors"
import "sync"

type Item struct {
	value int
	next *Item // link to the next Post
}

var MAX = 1000

type stack interface {
    Size() int
	Push(item Item) error
	Pop() Item
	NewStack() *StackImpl
	PrintStack()
}

type StackImpl struct {
	front *Item
	size int
	stackLock  sync.RWMutex
	empty sync.Cond
	full sync.Cond
}

func NewStack() *StackImpl {
	stack := new(StackImpl)
	stack.stackLock = *new(sync.RWMutex)
	return stack
}

func (stack *StackImpl) Size() int{
	return stack.size
}

func (stack *StackImpl) push(item *Item) error {
	if item == nil {
		return errors.New("Item is nil")
	}
	//  ThreadSafety
	stack.stackLock.Lock()
	if stack.size == MAX {
		return errors.New("Stack is full");
	}
    item.next = stack.front;
	stack.front = item;
	stack.size++
	defer stack.stackLock.Unlock()
	return nil
}


func (stack *StackImpl) pop() *Item {
	//  ThreadSafety
	stack.stackLock.Lock()
	if stack.size == 0 {
		return nil;
	}
	item := stack.front
	stack.front = stack.front.next;
	stack.size--;
	defer stack.stackLock.Unlock()
	return item
}


func (stack *StackImpl) PrintStackContent()  {
	item := stack.front;
	for {
        if item == nil {
			break
		} else {
			print("--")
			print(item.value)
			item = item.next
		}
	}
}