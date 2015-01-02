package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else if v == lim {
		fmt.Printf("%g = %g\n", v, lim)
	} else {
		fmt.Printf("%g > %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(
		pow(2, 2, 4),
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
