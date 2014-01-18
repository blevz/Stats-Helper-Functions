package main

import (
	"math"
)

func makeBernoulliDistribution(p probability) (rv randomVarStore) {
	rv.init()
	rv.data[0] = 1 - p
	rv.data[1] = p
	return
}

func makeBinomialDistribution(p probability, numTrials int) (rv randomVarStore) {
	rv.init()
	for x := 0; x <= numTrials; x++ {
		cpart := float64(Choose(numTrials, x))
		pex := math.Pow(float64(p), float64(x))
		qex := math.Pow(float64(1-p), float64(numTrials-x))
		rv.data[space_t(x)] = probability(cpart * pex * qex)

	}
	return
}

func makeHyperGeomDistribution(n1, n2, nToPick int) (rv randomVarStore) {
	rv.init()
	for x := 0; x <= nToPick; x++ {
		numer := Choose(n1, x) * Choose(n2, nToPick-x)
		denom := Choose(n1+n2, nToPick)
		rv.data[space_t(x)] = probability(numer / denom)
	}
	return
}
