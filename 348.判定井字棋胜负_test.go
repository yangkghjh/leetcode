package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTicTacToe(t *testing.T) {
	Convey("[348] 判定井字棋胜负", t, func() {
		t := TicTacToeConstructor(3)
		So(t.Move(0, 0, 1), ShouldEqual, 0)
		So(t.Move(0, 1, 1), ShouldEqual, 0)
		So(t.Move(0, 2, 1), ShouldEqual, 1)

		t = TicTacToeConstructor(3)
		So(t.Move(0, 0, 2), ShouldEqual, 0)
		So(t.Move(0, 1, 2), ShouldEqual, 0)
		So(t.Move(0, 2, 2), ShouldEqual, 2)

		t = TicTacToeConstructor(3)
		So(t.Move(0, 0, 1), ShouldEqual, 0)
		So(t.Move(1, 1, 2), ShouldEqual, 0)
		So(t.Move(0, 1, 1), ShouldEqual, 0)
		So(t.Move(0, 2, 2), ShouldEqual, 0)
		So(t.Move(2, 0, 1), ShouldEqual, 0)
		So(t.Move(1, 0, 2), ShouldEqual, 0)
		So(t.Move(2, 1, 1), ShouldEqual, 0)
	})
}
