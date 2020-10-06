package leetcode

import (
	"testing"
)

func TestMinStackConstructor(t *testing.T) {
	obj := MinStackConstructor()
	obj.Push(1)
	obj.Pop()
	obj.Top()
	obj.GetMin()
}
