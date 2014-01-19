package main

const (
	With_Replacement_Order_Matters = iota
	Without_Replacement_Order_Matters
	With_Replacement_Order_Not_Matters
	Without_Replacement_Order_Not_Matters
)

func selectNFromR(n, r, typeOfSelect int) int {
	switch typeOfSelect {
	case With_Replacement_Order_Matters:
		toRet := 1
		for x := 0; x < r; x++ {
			toRet *= n
		}
		return toRet
	case With_Replacement_Order_Not_Matters:
		toRet := 1
		for x := n; x <= n-1+r; x++ {
			toRet *= x
		}
		toRet /= Fact(r)
		return toRet

	case Without_Replacement_Order_Matters:
		return Perm(n, r)

	case Without_Replacement_Order_Not_Matters:
		return Choose(n, r)
	}
	return 5
}

func selectFromTwoPop(v1, v2, n, x int) float64 {
	num := Choose(v1, x) * Choose(v2, n-x)
	var toReturn float64 = float64(num) / float64(Choose(v1+v2, n))
	return toReturn
}
