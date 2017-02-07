package factorial

func Solution(n uint64) uint64 {
	//return Loop(n)
	//return Recursion(n)
	return TailRecursion(n, 1)
}

func Loop(n uint64) (x uint64) {
	x = 1
	for i := n; i > 0; i-- {
		x *= i
	}
	return
}

func Recursion(n uint64) uint64 {
	if n > 0 {
		return n * Recursion(n-1)
	}
	return 1
}

func TailRecursion(n, t uint64) uint64 {
	if n > 0 {
		return TailRecursion(n-1, n*t)
	}
	return t
}
