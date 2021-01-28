package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMergeTwoArray(t *testing.T) {
	Convey("[88] 合并两个有序数组", t, func() {
		arraym := [][]int{
			[]int{1, 2, 3, 0, 0, 0},
			[]int{0},
			[]int{2, 0},
		}
		m := []int{3, 0, 1}
		arrayn := [][]int{
			[]int{2, 5, 6},
			[]int{1},
			[]int{1},
		}
		n := []int{3, 1, 1}

		expected := [][]int{
			[]int{1, 2, 2, 3, 5, 6},
			[]int{1},
			[]int{1, 2},
		}

		for i := range arraym {
			res := arraym[i]
			merge(res, m[i], arrayn[i], n[i])
			So(res, ShouldResemble, expected[i])
		}
	})
}
