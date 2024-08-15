package maths

func Iterative(n int) int {
	a, b := 0, 1
	if n == 0 || n == 1 {
		return n
	}

	for i := 2; i <= n; i++ {
		c := a + b
		a, b = b, c
	}
	return b
}
