package main

import (
	"fmt"
	"math"
)

type randomVar struct {
	space []space_t
	f     func(space_t) probability
}

type randomVarStore struct {
	data map[space_t]probability
}

func (rv randomVar) toStore() (r randomVarStore) {
	r.init()
	for _, val := range rv.space {
		fmt.Println("adding ", val)
		r.data[val] = rv.f(val)
	}
	return
}

func (r *randomVarStore) init() {
	r.data = make(map[space_t]probability)
}

func (r randomVar) expectedValue() probability {
	var toReturn probability = 0.0
	return toReturn
}

func (r randomVarStore) expectedValue() probability {
	var toReturn probability = 0
	for u, p := range r.data {
		toReturn += probability(u) * p
	}
	return toReturn
}

func (r randomVarStore) meanVariance() probability {
	var toReturn probability = 0
	expectedValue := r.expectedValue()
	for u, p := range r.data {
		z := (probability(u) - expectedValue)
		toReturn += p * z * z
	}
	return toReturn
}

func (r randomVarStore) stdDeviation() probability {
	//Is there a better way to do this?
	return probability(math.Sqrt(float64(r.meanVariance())))
}

func (r randomVarStore) print() {
	for u, p := range r.data {
		fmt.Printf("P(X=%f)=%f, ", u, p)
	}
	fmt.Println("")
}

func (r randomVarStore) sumProbs() probability {
	sum := probability(0)
	for _, p := range r.data {
		sum += p
	}
	return sum
}
