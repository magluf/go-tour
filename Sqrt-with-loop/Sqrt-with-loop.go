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
	var list [1024]float64
	for i := 1; i <= 1024; i++ {
		list[i-1] = Sqrt(float64(i))
	}
	for i := 0; i < len(list); i++ {
		if list[i] == float64(int64(list[i])) {
			fmt.Println(math.Round(list[i]*list[i]), "->", list[i], "\n")
		}
	}
}
