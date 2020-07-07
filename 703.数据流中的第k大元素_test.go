package leetcode

import (
	"container/heap"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestKthLargest_Add(t *testing.T) {
	k := NewKthLargest(3, []int{4, 5, 8, 2})
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields KthLargest
		args   args
		want   int
	}{
		{
			name:   "1",
			fields: k,
			args: args{
				val: 3,
			},
			want: 4,
		},
		{
			name:   "2",
			fields: k,
			args: args{
				val: 5,
			},
			want: 5,
		},
		{
			name:   "3",
			fields: k,
			args: args{
				val: 10,
			},
			want: 5,
		},
		{
			name:   "4",
			fields: k,
			args: args{
				val: 9,
			},
			want: 8,
		},
		{
			name:   "5",
			fields: k,
			args: args{
				val: 4,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.Add(tt.args.val); got != tt.want {
				t.Errorf("KthLargest.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSmallTopHeap(t *testing.T) {
	Convey("小顶堆", t, func() {
		h := new(SmallTopHeap)

		heap.Push(h, 1)
		So(h.Len(), ShouldEqual, 1)
	})
}
