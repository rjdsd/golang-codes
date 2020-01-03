package  blockingstack
import "sync"
//import "fmt"


var MAX = 1000


type Item struct {
	value int
	next *Item // link to the next Post
}

type blockingstack interface {
    Size() int
	Push(item Item) 
	Pop() Item
	NewBlockingStack() *BlockingStackImpl
	PrintBlockingStack()
}

type BlockingStackImpl struct {
	front *Item
	size int
	stackLock  sync.RWMutex
	empty sync.Cond
	full sync.Cond
}

func NewBlockingStack() *BlockingStackImpl {
	stack := new(BlockingStackImpl)
	stack.stackLock = *new(sync.RWMutex)
	return stack
}

func (stack *BlockingStackImpl) Size() int{
	return stack.size
}

func (stack *BlockingStackImpl) Push(item *Item)  {
	//  to ensure only 1 can enStack at one time
	for {
		stack.stackLock.Lock()
        if stack.size < MAX {
			break;
		}
		// self preemption : necessary step for Deadlock avoidance
		stack.stackLock.Unlock()
	}
	item.next = stack.front;
	stack.front = item;
	stack.size++
	stack.stackLock.Unlock()
}


func (stack *BlockingStackImpl) Pop() *Item  {
	//  to ensure only 1 can pop at one time
	for {
		stack.stackLock.Lock()
        if stack.size > 0 {
			break
		}
		// self preemption : necessary step for Deadlock avoidance
		stack.stackLock.Unlock()
	}
	item := stack.front
	stack.front = stack.front.next
	stack.size--
	defer stack.stackLock.Unlock()
	return item
}

func (stack *BlockingStackImpl) PrintStack()  {
	//  to ensure only 1 can deStack at one time
	item := stack.front
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