package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	count := 100

	for count > 0 {
		count -= 1
		chg := (z*z - x) / (2 * z)
		fmt.Println(z)
		if math.Abs(chg) < 0.000000000000001 {
			break
		}
		z -= chg
	}

	return z
}

func main() {
	fmt.Println(Sqrt(4))
}
