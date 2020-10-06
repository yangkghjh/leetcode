package offer

import (
	"reflect"
	"testing"
)

func Test_levelOrder3(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				root: NewBinaryTree(3, 9, 20, null, null, 15, 7),
			},
			want: [][]int{
				[]int{3},
				[]int{20, 9},
				[]int{15, 7},
			},
		},
		{
			name: "2",
			args: args{
				root: NewBinaryTree(1, 2, 3, 4, null, null, 5),
			},
			want: [][]int{
				[]int{1},
				[]int{3, 2},
				[]int{4, 5},
			},
		},
		{
			name: "3",
			args: args{
				root: NewBinaryTree(0, 2, 4, 1, null, 3, -1, 5, 1, null, 6, null, 8),
			},
			want: [][]int{
				[]int{0},
				[]int{4, 2},
				[]int{1, 3, -1},
				[]int{8, 6, 1, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder3(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder3() = %v, want %v", got, tt.want)
			}
		})
	}
}
