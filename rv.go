package main

import (
	"fmt"
	"math"
	"sort"
)

type randomVar struct {
	space []float64
	f     func(float64) float64
}

type randomVarStore struct {
	data map[float64]float64
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
	r.data = make(map[float64]float64)
}

func (rv randomVarStore) getSortedKeys() (toRet []float64) {
	toRet = make([]float64, 0)
	for k, _ := range rv.data {
		toRet = append(toRet, k)
	}
	sort.Float64s(toRet)
	return
}

func (r randomVar) expectedValue() float64 {
	var toReturn float64 = 0.0
	return toReturn
}

func (r randomVarStore) expectedValue() float64 {
	var toReturn float64 = 0
	for u, p := range r.data {
		toReturn += float64(u) * p
	}
	return toReturn
}

func (r randomVarStore) meanVariance() float64 {
	var toReturn float64 = 0
	expectedValue := r.expectedValue()
	for u, p := range r.data {
		z := (float64(u) - expectedValue)
		toReturn += p * z * z
	}
	return toReturn
}

func (r randomVarStore) stdDeviation() float64 {
	//Is there a better way to do this?
	return float64(math.Sqrt(float64(r.meanVariance())))
}

func (r randomVarStore) print() {
	for u, p := range r.data {
		fmt.Printf("P(X=%f)=%f, ", u, p)
	}
	fmt.Printf("\n")
}

func (r randomVarStore) sumProbs() float64 {
	sum := float64(0)
	for _, p := range r.data {
		sum += p
	}
	return sum
}

func (r *randomVarStore) convertToCDF() {
	for i, val := range r.data {
		r.data[i] = r.data[i-1] + val
	}
}
