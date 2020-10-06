package offer

import "testing"

func Test_verifyPostorder(t *testing.T) {
	type args struct {
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				postorder: []int{1, 3, 2, 6, 5},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				postorder: []int{1, 6, 3, 2, 5},
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				postorder: []int{5, 4, 3, 2, 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verifyPostorder(tt.args.postorder); got != tt.want {
				t.Errorf("verifyPostorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
