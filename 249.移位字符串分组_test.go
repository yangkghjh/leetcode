package leetcode

import (
	"reflect"
	"testing"
)

func Test_groupStrings(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "1",
			args: args{
				strings: []string{"az", "ba"},
			},
			want: [][]string{
				{"az", "ba"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupStrings(tt.args.strings); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
