package ds
import "errors"
import "sync"

type blockingstack interface {
    Size() int
	Push(item Item) error
	Pop() Item
	NewStack() *BlockingStackImpl
	PrintStack()
}

type BlockingStackImpl struct {
	front *Item
	size int
	stackLock  sync.RWMutex
	empty sync.Cond
	full sync.Cond
}

func NewStack() *BlockingStackImpl {
	stack := new(BlockingStackImpl)
	stack.stackLock = *new(sync.RWMutex)
	return stack
}

func (stack *BlockingStackImpl) Size() int{
	return stack.size
}

func (stack *BlockingStackImpl) push(item *Item) error {
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


func (stack *BlockingStackImpl) pop() *Item {
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


func (stack *BlockingStackImpl) PrintStackContent()  {
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