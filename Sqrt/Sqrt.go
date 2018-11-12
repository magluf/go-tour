package main

import (
	"fmt"
	"math"
)

// Sqrt frescura
func Sqrt(x float64) float64 {
	z := 1.0
	k := 0.0
	for {
		z -= (z*z - x) / (2 * z)
		if math.Floor(z*10e10)/10e10 == k {
			break
		}
		k = math.Floor(z*10e10) / 10e10
	}
	return z
}

func main() {
	fmt.Println(Sqrt(5))
}
