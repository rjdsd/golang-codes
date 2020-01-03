package blockingstack



import (
    "testing"
    "github.com/stretchr/testify/assert"
    "fmt"
  )

func TestQueue(t *testing.T) {
    // assert equality
    assert.Equal(t, 123, 123, "they should be equal")
    
    stack := NewBlockingStack()
    var sum2,sum3,iter int
    iter = 10000
    for i := 1; i <= iter; i++ {
        item := Item{100,nil}
        sum2 += 100
        go stack.Push(&item)
    } 

    
    for i := 0; i < iter; i++ {
        sum3 = sum3 + (*stack.Pop()).value
	}  
    
    assert.Equal(t, sum2, sum3, "they should be equal")
    fmt.Println("Sum3")
    fmt.Println(sum3)
    
}