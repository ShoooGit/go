package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	i := 1
	z := 1.0

	for i < 10 {
		z -= (z*z - x) / (2*z)
		//fmt.Println(z)
		i += 1
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}