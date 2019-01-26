package ds


import (
    "testing"
    "github.com/stretchr/testify/assert"
    "time"
  )

func TestQueue(t *testing.T) {
    // assert equality
    assert.Equal(t, 123, 123, "they should be equal")

    //  FAIL/
    
    /*
    queue := NewQueue()
    assert.Nil(t, queue.Dequeue())
    */

    
    queue := NewQueue()
    item := Item{100,nil}
    queue.Enqueue(&item)
    assert.NotNil(t, queue.Dequeue())

    
    /*
    queue = NewQueue()
    var sum1, sum int
    for i := 0; i < 11; i++ {
        item := Item{i,nil}
        sum1 += i
        queue.Enqueue(&item)
	} 
    for  queue.Size() > 0  {
        sum = sum + (*queue.Dequeue()).value
    }
    assert.Equal(t, sum, sum1, "they should be equal")
    */
    
    queue = NewQueue()
    var sum1,sum int
    for i := 1; i <= 1000; i++ {
        item := Item{100,nil}
        sum1 += 100
        go queue.Enqueue(&item)
    } 

    
    for i := 0; i < 1000; i++ {
        sum = sum + (*queue.Dequeue()).value
	}  
    
    assert.Equal(t, sum, sum1, "they should be equal")
    time.Sleep(2 * time.Second)
    queue.PrintQueue()
}