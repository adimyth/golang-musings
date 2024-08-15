package maths

func Recursive(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return Recursive(n-1) + Recursive(n-2)
}
