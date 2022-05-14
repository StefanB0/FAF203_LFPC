package lib

import "math"

func Sieve_is_prime(n int64) bool {
	c:= make([]bool, n+1)
	for i := range c {
		c[i] = true
	}

	c[0] = false
	c[1] = false

	lowerLimit := int64(math.Sqrt(float64(n)))
	for i := int64(2); i <= lowerLimit; i++{
		if c[i] {
			j:= i*i
			for j <= n {
				c[j] = false
				j+= i
			}
		}
	}

	return c[n];
}
