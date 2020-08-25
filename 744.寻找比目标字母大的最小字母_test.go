package leetcode

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{name: "1", args: args{letters: []byte{'c', 'f', 'j'}, target: 'a'}, want: 'c'},
		{name: "2", args: args{letters: []byte{'a', 'b'}, target: 'z'}, want: 'a'},
		{name: "3", args: args{letters: []byte{'c', 'f', 'j'}, target: 'c'}, want: 'f'},
		{name: "4", args: args{letters: []byte{'c', 'f', 'j'}, target: 'd'}, want: 'f'},
		{name: "5", args: args{letters: []byte{'c', 'f', 'j'}, target: 'g'}, want: 'j'},
		{name: "6", args: args{letters: []byte{'c', 'f', 'j'}, target: 'j'}, want: 'c'},
		{name: "7", args: args{letters: []byte{'c', 'f', 'j'}, target: 'k'}, want: 'c'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
