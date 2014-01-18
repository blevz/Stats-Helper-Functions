package main

func Choose(a, b int) int {
	if a-b > b {
		b = a - b
	}
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

func ChooseB(a, b int64) int64 {
	if a-b > b {
		b = a - b
	}
	var r = PermB(a, b) / FactB(a-b)
	return r
}

func PermB(a, b int64) int64 {
	r := int64(1)
	for x := a; x > b; x-- {
		r *= x
	}
	return r
}

func FactB(a int64) int64 {
	r := int64(1)
	for x := int64(2); x <= a; x++ {
		r *= x
	}
	return r
}
