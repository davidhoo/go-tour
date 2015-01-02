package main

import "fmt"

func main() {
	sum := 1
	for sum < 10 {
		fmt.Println(sum, "======")
		sum += sum
	}
	fmt.Println(sum)
}
