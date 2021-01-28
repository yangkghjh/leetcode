package leetcode

import (
	"testing"
)

func TestNewMyQueue(t *testing.T) {
	obj := NewMyQueue()
	obj.Push(1)
	obj.Pop()
	obj.Peek()
	obj.Empty()
}
