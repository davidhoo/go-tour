package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	matrix := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		matrix[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			matrix[i][j] = uint8((i ^ j) * (i ^ j))
		}
	}

	return matrix
}

func main() {
	pic.Show(Pic)
}
