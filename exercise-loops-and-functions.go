package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := 1.0
    for i := 0; i < 10; i++ {
        z = z - (z*z-x)/(z*2)
        fmt.Println(">>>>>>>", z)
    }
    return z
}

func main() {
    fmt.Println("my func:", Sqrt(2))
    fmt.Println("math.Sqrt:", math.Sqrt(2))
}
