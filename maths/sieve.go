package maths

func Sieve(n int) []int {
	// Initialize an array of length n & set all of them to True
	primes := make([]bool, n)
	for i := range primes {
		primes[i] = true
	}

	for i := 2; i*i < n; i++ {
		if primes[i] {
			for j := i; j*i < n; j++ {
				primes[i*j] = false
			}
		}
	}

	prime_numbers := []int{}
	for i := 0; i < n; i++ {
		if primes[i] == true {
			prime_numbers = append(prime_numbers, i)
		}
	}

	return prime_numbers
}
