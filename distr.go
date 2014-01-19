package main

import (
	"math"
)

func makeUniformDistribution(n int) (rv randomVarStore) {
	rv.init()
	for x := 0; x < n; x++ {
		rv.data[float64(x)] = float64(1.0 / float64(n))
	}
	return
}

func makeBernoulliDistribution(p float64) (rv randomVarStore) {
	rv.init()
	rv.data[0] = 1 - p
	rv.data[1] = p
	return
}

func makeBinomialDistribution(p float64, numTrials int) (rv randomVarStore) {
	rv.init()
	for x := 0; x <= numTrials; x++ {
		cpart := float64(Choose(numTrials, x))
		pex := math.Pow(float64(p), float64(x))
		qex := math.Pow(float64(1-p), float64(numTrials-x))
		rv.data[float64(x)] = float64(cpart * pex * qex)

	}
	return
}

func makeHyperGeomDistribution(n1, n2, nToPick int) (rv randomVarStore) {
	rv.init()
	for x := 0; x <= nToPick; x++ {
		numer := float64(Choose(n1, x) * Choose(n2, nToPick-x))
		denom := float64(Choose(n1+n2, nToPick))
		rv.data[float64(x)] = float64(numer / denom)
	}
	return
}

const (
	smallEp = .000000001
)

func makePoissonDistribution(param float64) (rv randomVarStore) {
	rv.init()
	for x := 0; ; x++ {
		numer := math.Pow(param, float64(x)) * math.Pow(math.E, -param)
		denom := float64(Fact(x))
		rv.data[float64(x)] = float64(numer / denom)
		if rv.data[float64(x)] < smallEp {
			break
		}
	}
	return
}
