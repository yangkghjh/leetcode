package offer

import "testing"

func Test_validateStackSequences(t *testing.T) {
	type args struct {
		pushed []int
		popped []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				pushed: []int{1, 2, 3},
				popped: []int{3, 2, 1},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				pushed: []int{1, 2, 3},
				popped: []int{3, 2},
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				pushed: []int{1, 2, 3},
				popped: []int{3, 3, 3},
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				pushed: []int{},
				popped: []int{},
			},
			want: true,
		},
		{
			name: "5",
			args: args{
				pushed: []int{1, 2, 3, 4, 5},
				popped: []int{4, 5, 3, 2, 1},
			},
			want: true,
		},
		{
			name: "6",
			args: args{
				pushed: []int{1, 2, 3, 4, 5},
				popped: []int{4, 3, 5, 1, 2},
			},
			want: false,
		},
		{
			name: "7",
			args: args{
				pushed: []int{2, 1, 0},
				popped: []int{1, 2, 0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateStackSequences(tt.args.pushed, tt.args.popped); got != tt.want {
				t.Errorf("validateStackSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
