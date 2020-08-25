package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMyPow(t *testing.T) {
	Convey("[50] Pow(x, n)", t, func() {
		xs := []float64{1, 2, 2, -2, 2, 2.1, 2, 0.44528, 1.72777}
		ns := []int{1, 2, -1, 3, 10, 3, -2, 0, 7}
		expected := []float64{
			1, 4, 0.5, -8, 1024, 9.261, 0.25, 1, 45.96227345465138,
		}

		for i, x := range xs {
			So(myPow(x, ns[i]), ShouldAlmostEqual, expected[i])
		}
	})
}
