package ds

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "time"
)

func TestStack(t *testing.T) {
    // assert equality
    assert.Equal(t, 123, 123, "they should be equal")


    stack := NewStack()
    item := Item{100,nil}
    stack.push(&item)
    assert.NotNil(t, stack.pop())


    stack = NewStack()
    var sum1,sum int
    for i := 1; i <= 1000; i++ {
        item := Item{100,nil}
        sum1 += 100
        go stack.push(&item)
    }
    time.Sleep(2 * time.Second)
    for i := 0; i < 1000; i++ {
        sum = sum + (*stack.pop()).value
	}
    assert.Equal(t, sum, sum1, "they should be equal")


    stack = NewStack()
    assert.Error(t, stack.push(nil), "Item is nil");

}