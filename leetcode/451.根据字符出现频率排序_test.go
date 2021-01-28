package leetcode

import "testing"

func Test_frequencySort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "trreee",
			},
			want: "eeerrt",
		},
		{
			name: "2",
			args: args{
				s: "ccccaaa",
			},
			want: "ccccaaa",
		},
		{
			name: "3",
			args: args{
				s: "AAbbba",
			},
			want: "bbbAAa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := frequencySort(tt.args.s); got != tt.want {
				t.Errorf("frequencySort() = %v, want %v", got, tt.want)
			}
		})
	}
}
