package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Rv(t *testing.T) {
	var a randomVarStore
	a.data = make(map[space_t]probability)
	a.data[0] = .2
	a.data[2] = .2
	a.data[4] = .2
	a.data[6] = .2
	a.data[8] = .2

	var b randomVarStore
	b.data = make(map[space_t]probability)
	b.data[0] = .2
	b.data[1] = .2
	b.data[2] = .2
	b.data[3] = .2
	b.data[4] = .2

	Convey("the sum of probability should always be near 1", t, func() {
		So(1.0, ShouldAlmostEqual, a.sumProbs())
		So(1.0, ShouldAlmostEqual, b.sumProbs())

	})

	Convey("Test means", t, func() {
		So(4, ShouldEqual, a.expectedValue())
		So(2, ShouldEqual, b.expectedValue())
	})

	Convey("Test varian", t, func() {
		So(8, ShouldEqual, a.meanVariance())
		So(2, ShouldEqual, b.expectedValue())
	})

}
