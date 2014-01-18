package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCombination(t *testing.T) {

	Convey("Factorial Test", t, func() {
		So(2, ShouldEqual, Fact(2))
		So(6, ShouldEqual, Fact(3))
		So(24, ShouldEqual, Fact(4))
		So(120, ShouldEqual, Fact(5))
	})

	Convey("Choose Test", t, func() {
		So(28, ShouldEqual, Choose(8, 2))
		So(21, ShouldEqual, Choose(7, 2))
		So(15, ShouldEqual, Choose(6, 2))
		So(10, ShouldEqual, Choose(5, 2))
		So(6, ShouldEqual, Choose(4, 2))
		So(10, ShouldEqual, Choose(5, 3))

	})

	Convey("Perm Test", t, func() {
		So(20, ShouldEqual, Perm(5, 3))
		So(120, ShouldEqual, Perm(6, 3))
	})
}
