package offer

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "1",
			args: args{
				n: 1,
			},
			want: []float64{0.16667, 0.16667, 0.16667, 0.16667, 0.16667, 0.16667},
		},
		{
			name: "2",
			args: args{
				n: 2,
			},
			want: []float64{0.02778, 0.05556, 0.08333, 0.11111, 0.13889, 0.16667, 0.13889, 0.11111, 0.08333, 0.05556, 0.02778},
		},
		{
			name: "3",
			args: args{
				n: 3,
			},
			want: []float64{0.02778, 0.05556, 0.08333, 0.11111, 0.13889, 0.16667, 0.13889, 0.11111, 0.08333, 0.05556, 0.02778},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// got := twoSum(tt.args.n)
			// t.Errorf("twoSum() = %v, want %v", got, tt.want)
		})
	}
}
