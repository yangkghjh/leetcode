package offer

import (
	"testing"
)

func Test_mirrorTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "1",
			args: args{
				root: NewBinaryTree(4, 2, 7, 1, 3, 6, 9),
			},
			want: NewBinaryTree(4, 7, 2, 9, 6, 3, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mirrorTree(tt.args.root); !TreeNodeDeepEqual(got, tt.want) {
				t.Errorf("mirrorTree() = %v, want %v", SerializeBinaryTree(got), SerializeBinaryTree(tt.want))
			}
		})
	}
}
