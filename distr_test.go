package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Distributions(t *testing.T) {
	a := makeBernoulliDistribution(.5)
	b := makeBinomialDistribution(.5, 8)
	c := makeUniformDistribution(6)
	d := makeHyperGeomDistribution(5, 5, 3)

	Convey("Probability should always sum to around 1", t, func() {

		So(1.0, ShouldAlmostEqual, a.sumProbs())
		So(1.0, ShouldAlmostEqual, b.sumProbs())
		So(1.0, ShouldAlmostEqual, c.sumProbs())
		So(1.0, ShouldAlmostEqual, d.sumProbs())

	})
}
