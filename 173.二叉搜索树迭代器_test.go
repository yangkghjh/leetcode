package leetcode

import "testing"

func TestBSTIterator_Next(t *testing.T) {
	b1 := NewBSTIterator(NewBinaryTree(7, 3, 15, -1, -1, 9, 20))

	tests := []struct {
		name string
		this *BSTIterator
		want []int
	}{
		{
			name: "1",
			this: &b1,
			want: []int{3, 7, 9, 15, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, i := range tt.want {
				if !tt.this.HasNext() {
					t.Errorf("BSTIterator.HasNext() = %v, want %v", false, true)
					break
				}
				if got := tt.this.Next(); got != i {
					t.Errorf("BSTIterator.Next() = %v, want %v", got, tt.want)
					break
				}
			}

			if tt.this.HasNext() {
				t.Errorf("BSTIterator.HasNext() = %v, want %v", true, false)
			}
		})
	}
}
