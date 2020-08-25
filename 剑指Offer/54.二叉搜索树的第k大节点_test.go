package offer

import "testing"

func Test_kthLargest(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				root: NewBinaryTree(3, 1, 4, null, 2),
				k:    1,
			},
			want: 4,
		},
		{
			name: "2",
			args: args{
				root: NewBinaryTree(5, 3, 6, 2, 4, null, null, 1),
				k:    3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
