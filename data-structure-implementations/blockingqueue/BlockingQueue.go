package blockingqueue
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
	//  Ensure thread safety
	for {
		queue.queueLock.Lock()
        if queue.size < MAX {
			break;
		}
		//Release lock voluntarily to avoid Deadlock and try to aquire lock again
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
	queue.queueLock.Unlock()
}


func (queue *BlockingQueueImpl) Dequeue() *Item  {
	//  Ensure thread safety
	for {
		queue.queueLock.Lock()
        if queue.size > 0 {
			break
		}
		queue.queueLock.Unlock()
	}
	item := queue.front
	queue.front = queue.front.next
	queue.size--
	//Release lock voluntarily to avoid Deadlock and try to aquire lock again
	defer queue.queueLock.Unlock()
	return item
}

func (queue *BlockingQueueImpl) PrintQueue()  {
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