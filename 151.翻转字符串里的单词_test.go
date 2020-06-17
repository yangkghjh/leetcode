package leetcode

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "the sky is blue",
			args: args{
				s: "the sky is blue",
			},
			want: "blue is sky the",
		},
		{
			name: "",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "  ",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "  hello world!  ",
			args: args{
				s: "  hello world!  ",
			},
			want: "world! hello",
		},
		{
			name: "a good   example",
			args: args{
				s: "a good   example",
			},
			want: "example good a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
