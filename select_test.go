package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Select(t *testing.T) {

	Convey("replacement order matters", t, func() {
		a := selectNFromR(5, 2, With_Replacement_Order_Matters)
		b := selectNFromR(5, 3, With_Replacement_Order_Matters)

		So(a, ShouldEqual, 25)
		So(b, ShouldEqual, 125)
	})

	Convey("replacement order doesn't matters", t, func() {
		a := selectNFromR(5, 2, With_Replacement_Order_Not_Matters)
		b := selectNFromR(5, 3, With_Replacement_Order_Not_Matters)

		So(a, ShouldEqual, 15)
		So(b, ShouldEqual, 35)
	})

	Convey("No replacement order does matters", t, func() {
		a := selectNFromR(5, 2, Without_Replacement_Order_Matters)
		b := selectNFromR(5, 3, Without_Replacement_Order_Matters)

		So(a, ShouldEqual, 60)
		So(b, ShouldEqual, 20)
	})

	Convey("No replacement order doesn't matters", t, func() {
		a := selectNFromR(5, 2, Without_Replacement_Order_Not_Matters)
		b := selectNFromR(5, 3, Without_Replacement_Order_Not_Matters)

		So(a, ShouldEqual, 10)
		So(b, ShouldEqual, 10)
	})

}
