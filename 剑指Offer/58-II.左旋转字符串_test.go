package offer

import "testing"

func Test_reverseLeftWords(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "abcedf",
				n: 2,
			},
			want: "cedfab",
		},
		{
			name: "2",
			args: args{
				s: "abcedf",
				n: 0,
			},
			want: "abcedf",
		},
		{
			name: "3",
			args: args{
				s: "abcedf",
				n: 8,
			},
			want: "cedfab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseLeftWords(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("reverseLeftWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
