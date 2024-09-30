package main

import "fmt"

// Persistence function to calculate multiplicative persistence
func Persistence(n int) int {
	count := 0

	// Continue multiplying digits until n is reduced to a single digit
	for n >= 10 {
		n = multiplyDigits(n)
		count++
	}

	return count
}

// Helper function to multiply digits of a number
func multiplyDigits(n int) int {
	product := 1
	for n > 0 {
		product *= n % 10
		n /= 10
	}
	return product
}

func main() {
	fmt.Println(Persistence(39))  // Output: 3
	fmt.Println(Persistence(999)) // Output: 4
	fmt.Println(Persistence(4))   // Output: 0
}
