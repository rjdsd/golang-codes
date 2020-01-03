package ds
import "sync"
//import "fmt"


var MAX = 1000


type Item struct {
	value int
	next *Item // link to the next Post
}

type blockingqueue interface {
    Size() int
	Enqueue(item Item) 
	Dequeue() Item
	NewQueue() *BlockingQueueImpl
	PrintQueue()
}

type BlockingQueueImpl struct {
	front *Item
	end *Item
	size int
	queueLock  sync.RWMutex
	empty sync.Cond
	full sync.Cond
}

func NewQueue() *BlockingQueueImpl {
	queue := new(BlockingQueueImpl)
	queue.queueLock = *new(sync.RWMutex)
	return queue
}

func (queue *BlockingQueueImpl) Size() int{
	return queue.size
}

func (queue *BlockingQueueImpl) Enqueue(item *Item)  {
	//  to ensure only 1 can enqueue at one time
	for {
		queue.queueLock.Lock()
		//fmt.Println("enque lock")
        if queue.size < MAX {
			break;
		}
		//fmt.Println("enque unlock")
		queue.queueLock.Unlock()
	}
	if queue.end != nil {
		(queue.end).next = item
	}
	queue.end = item
	if queue.front == nil {	
		queue.front = item
	}
	queue.size++
	//fmt.Println("enque unlock")
	queue.queueLock.Unlock()
}


func (queue *BlockingQueueImpl) Dequeue() *Item  {
	//  to ensure only 1 can dequeue at one time
	for {
		queue.queueLock.Lock()
		//fmt.Println("deque lock")
        if queue.size > 0 {
			break
		}
		queue.queueLock.Unlock()
		//fmt.Println("deque unlock")
	}
	item := queue.front
	queue.front = queue.front.next  //  getting a nil pointer at this line sometimes
	queue.size--
	//fmt.Println("deque unlock")
	defer queue.queueLock.Unlock()
	return item
}

func (queue *BlockingQueueImpl) PrintQueue()  {
	//  to ensure only 1 can dequeue at one time
	item := queue.front
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