package blockingqueue


import (
    "testing"
    "github.com/stretchr/testify/assert"
    "fmt"
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
    
    
    
    queue = NewQueue()
    var sum2,sum3,iter int
    iter = 50000
    for i := 1; i <= iter; i++ {
        item := Item{100,nil}
        sum2 += 100
        go queue.Enqueue(&item)
    } 

    
    for i := 0; i < iter; i++ {
        sum3 = sum3 + (*queue.Dequeue()).value
        fmt.Println("Sum3")
        fmt.Println(sum3)
        fmt.Println("Number of Element Dequeued")
        fmt.Println(i)
	}  
    
    assert.Equal(t, sum2, sum3, "they should be equal")

    
}