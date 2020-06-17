package leetcode

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "babad",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "cbbd",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "bb",
			args: args{
				s: "bb",
			},
			want: "bb",
		},
		{
			name: "b",
			args: args{
				s: "b",
			},
			want: "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
