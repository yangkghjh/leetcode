package offer

import "testing"

func Test_isStraight(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				nums: []int{0, 0, 1, 2, 5},
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				nums: []int{11, 10, 0, 0, 12},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStraight(tt.args.nums); got != tt.want {
				t.Errorf("isStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}
