package offer

import (
	"reflect"
	"testing"
)

func Test_levelOrder2(t *testing.T) {
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
				[]int{9, 20},
				[]int{15, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
