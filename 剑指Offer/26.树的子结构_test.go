package offer

import "testing"

func Test_isSubStructure(t *testing.T) {
	type args struct {
		A *TreeNode
		B *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				A: NewBinaryTree(1, 2, 3),
				B: NewBinaryTree(3, 1),
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				A: NewBinaryTree(3, 4, 5, 1, 2),
				B: NewBinaryTree(4, 1),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubStructure(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("isSubStructure() = %v, want %v", got, tt.want)
			}
		})
	}
}
