package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println("*p is:", *p)

	*p = 21
	fmt.Println("i is:", i)

	p = &j
	fmt.Println("*p is:", *p)

	*p = *p / 37
	fmt.Println("j is:", j)
}
