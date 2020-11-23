package offer

import "testing"

func Test_findNthDigit(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "10",
			args: args{
				n: 10,
			},
			want: 1,
		},
		{
			name: "11",
			args: args{
				n: 11,
			},
			want: 0,
		},
		{
			name: "20",
			args: args{
				n: 20,
			},
			want: 1,
		},
		{
			name: "199",
			args: args{
				n: 199,
			},
			want: 1,
		},
		{
			name: "200",
			args: args{
				n: 200,
			},
			want: 0,
		},
		{
			name: "201",
			args: args{
				n: 201,
			},
			want: 3,
		},
		{
			name: "1000",
			args: args{
				n: 1000,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNthDigit(tt.args.n); got != tt.want {
				t.Errorf("findNthDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
