package leetcode

import (
	"testing"
)

func TestNewMyStack(t *testing.T) {
	obj := NewMyStack()
	obj.Push(1)
	obj.Pop()
	obj.Top()
	obj.Empty()
}
