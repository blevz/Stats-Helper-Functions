package main

import (
	"fmt"
)

func Choose(a, b int) int {
	if a-b > b {
		b = a - b
	}
	var r = Perm(a, b) / Fact(a-b)
	if Verbosity >= Combinatorics_Verbosity {
		fmt.Printf("%d C %d = %d\n", a, b, r)
	}
	return r
}

func Perm(a, b int) int {
	var r int = 1
	for x := a; x > b; x-- {
		r *= x
	}
	if Verbosity >= Combinatorics_Verbosity {
		fmt.Printf("%d P %d = %d\n", a, b, r)
	}
	return r
}

func Fact(a int) int {
	var r int = 1
	for x := int(2); x <= a; x++ {
		r *= x
	}
	if Verbosity >= Combinatorics_Verbosity {
		fmt.Printf("%d! = %d\n", a, r)
	}
	return r
}
