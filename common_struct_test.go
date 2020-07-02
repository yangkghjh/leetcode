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

func TestTreeNodeDeepEqual(t *testing.T) {
	type args struct {
		m *TreeNode
		n *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				m: NewBinaryTree(1, 2, 3, 4, 5),
				n: NewBinaryTree(1, 2, 3, 4, 5),
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				m: NewBinaryTree(1, 2, 3, 4, 5),
				n: NewBinaryTree(1, 2, 3, 4, 2),
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				m: NewBinaryTree(1, 2, 3, 4, 5),
				n: NewBinaryTree(1, 2, 3, 4),
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				m: NewBinaryTree(1, 2, 3, 4, -1, 6, 7, 8),
				n: NewBinaryTree(1, 2, 3, 4, -1, 6, 7, 8, 9),
			},
			want: false,
		},
		{
			name: "5",
			args: args{
				m: NewBinaryTree(),
				n: NewBinaryTree(),
			},
			want: true,
		},
		{
			name: "6",
			args: args{
				m: NewBinaryTree(),
				n: NewBinaryTree(1),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TreeNodeDeepEqual(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("TreeNodeDeepEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
