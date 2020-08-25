package leetcode

import "testing"

func Test_replaceWords(t *testing.T) {
	type args struct {
		dict     []string
		sentence string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				dict:     []string{"cat", "bat", "rat"},
				sentence: "the cattle was rattled by the battery",
			},
			want: "the cat was rat by the bat",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceWords(tt.args.dict, tt.args.sentence); got != tt.want {
				t.Errorf("replaceWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
