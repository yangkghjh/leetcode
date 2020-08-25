package leetcode

import (
	"reflect"
	"testing"
)

func TestCodec_serialize(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		this *Codec
		args args
		want string
	}{
		{
			name: "1",
			this: new(Codec),
			args: args{
				root: NewBinaryTree(),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			if got := this.serialize(tt.args.root); got != tt.want {
				t.Errorf("Codec.serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodec_deserialize(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		this *Codec
		args args
		want *TreeNode
	}{
		{
			name: "1",
			this: new(Codec),
			args: args{
				data: "",
			},
			want: NewBinaryTree(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			if got := this.deserialize(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Codec.deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
