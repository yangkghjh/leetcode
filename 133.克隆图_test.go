package leetcode

import (
	"reflect"
	"testing"
)

func Test_cloneGraph(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "1",
			args: args{
				node: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cloneGraph(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cloneGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}
