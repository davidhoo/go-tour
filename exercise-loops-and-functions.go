package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		if math.Abs(z-(z-(z*z-x)/(z*2))) < 0.00000000000001 {
			return z
		} else {
			z = z - (z*z-x)/(z*2)
			fmt.Println(">>>>>z is:", z)
		}
	}
}

func main() {
	fmt.Println("  my func:", Sqrt(2))
	fmt.Println("math.Sqrt:", math.Sqrt(2))
}
