package leetcode

import (
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	type args struct {
		head *CopyRandomListNode
	}

	node4 := &CopyRandomListNode{Val: 1}
	node3 := &CopyRandomListNode{Val: 10, Next: node4}
	node2 := &CopyRandomListNode{Val: 11, Next: node3}
	node1 := &CopyRandomListNode{Val: 13, Next: node2}
	node0 := &CopyRandomListNode{Val: 7, Next: node1}
	node1.Random = node0
	node2.Random = node4
	node3.Random = node2
	node4.Random = node0

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				head: node0,
			},
			want: [][]int{{7, -1}, {13, 7}, {11, 1}, {10, 11}, {1, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := copyRandomList(tt.args.head)

			for _, v := range tt.want {
				if v[0] != got.Val {
					t.Errorf("copyRandomList() = %v, want %v", got.Val, v[0])
					break
				}
				if v[1] == -1 {
					if got.Random != nil {
						t.Errorf("copyRandomList() = %v, want %v", got.Random, v[1])
						break
					}
				} else {
					if v[1] != got.Random.Val {
						t.Errorf("copyRandomList() = %v, want %v", got.Random.Val, v[1])
						break
					}
				}
				got = got.Next
			}

		})
	}
}
