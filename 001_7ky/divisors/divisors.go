package main

import (
	"fmt"
)

func main() {
	var number int = 500000556
	fmt.Printf("Number %d have %d divisors\n", number, Divisors(number))
}

func Divisors(n int) int {
	//Put your code here
	var divisors int = 0
	for diviser := 1; diviser <= n; diviser++ {
		if n%diviser == 0 {
			divisors++
		}
	}
	return divisors
}
