package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewBinaryTree(t *testing.T) {
	Convey("NewBinaryTree", t, func() {
		root := NewBinaryTree(1, 2, 3)
		So(root.Val, ShouldEqual, 1)
		So(root.Left.Val, ShouldEqual, 2)
		So(root.Right.Val, ShouldEqual, 3)

		root = NewBinaryTree(5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, 1)
		So(root.Val, ShouldEqual, 5)
		So(root.Left.Val, ShouldEqual, 4)
		So(root.Right.Val, ShouldEqual, 8)

		So(root.Left.Left.Val, ShouldEqual, 11)
		So(root.Left.Right, ShouldEqual, nil)

		So(root.Right.Left.Val, ShouldEqual, 13)
		So(root.Right.Right.Val, ShouldEqual, 4)
	})
}
