package offer

import (
	"reflect"
	"testing"
)

func Test_getKthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}

	list1 := NewList(1, 2, 3, 4, 5)

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				head: list1,
				k:    2,
			},
			want: list1.Next.Next.Next,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKthFromEnd(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getKthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
