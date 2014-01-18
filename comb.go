package main

func Choose(a, b int) int {
	var r = Perm(a, b) / Fact(a-b)
	return r
}

func Perm(a, b int) int {
	var r int = 1
	for x := a; x > b; x-- {
		r *= x
	}
	return r
}

func Fact(a int) int {
	var r int = 1
	for x := int(2); x <= a; x++ {
		r *= x
	}
	return r
}
