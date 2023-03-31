package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		var r float64
		return r, ErrNegativeSqrt(x)
	}
	z := 1.0
	prev := 1.0
	i := 1
	for {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-prev) < 0.00000000000001 {
			return z, nil
		}
		prev = z
		i++
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
