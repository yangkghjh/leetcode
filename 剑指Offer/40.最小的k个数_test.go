package offer

import (
	"reflect"
	"testing"
)

func Test_getLeastNumbers(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				arr: []int{3, 2, 1},
				k:   2,
			},
			want: []int{2, 1},
		},
		{
			name: "2",
			args: args{
				arr: []int{0, 1, 2, 1},
				k:   1,
			},
			want: []int{0},
		},
		{
			name: "3",
			args: args{
				arr: []int{0, 1, 2, 1},
				k:   0,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLeastNumbers(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLeastNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
