package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

func main() {
	x, y := split(23)
	fmt.Println(y, x)
	fmt.Println(split(17))
}
