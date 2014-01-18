package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Distributions(t *testing.T) {

	a := makeBernoulliDistribution(.5)
	b := makeBernoulliDistribution(.001)

	//For large values of the second value the probability stops summing to 1
	c := makeBinomialDistribution(.5, 8)
	d := makeBinomialDistribution(.1, 10)

	e := makeUniformDistribution(6)
	f := makeUniformDistribution(1000)

	g := makeHyperGeomDistribution(5, 5, 3)
	h := makeHyperGeomDistribution(17, 10, 6)

	i := makePoissonDistribution(.9)
	j := makePoissonDistribution(.5)
	k := makePoissonDistribution(.1)

	Convey("Probability should always sum to around 1", t, func() {

		So(1.0, ShouldAlmostEqual, a.sumProbs())
		So(1.0, ShouldAlmostEqual, b.sumProbs())
		So(1.0, ShouldAlmostEqual, c.sumProbs())
		So(1.0, ShouldAlmostEqual, d.sumProbs())
		So(1.0, ShouldAlmostEqual, e.sumProbs())
		So(1.0, ShouldAlmostEqual, f.sumProbs())
		So(1.0, ShouldAlmostEqual, g.sumProbs())
		So(1.0, ShouldAlmostEqual, h.sumProbs())
		So(1.0, ShouldAlmostEqual, i.sumProbs())
		So(1.0, ShouldAlmostEqual, j.sumProbs())
		So(1.0, ShouldAlmostEqual, k.sumProbs())
	})
}
