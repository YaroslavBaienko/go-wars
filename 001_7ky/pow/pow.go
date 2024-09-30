package main

import (
	"fmt"
)

func main() {
	var n int = 5
	result := PowersOfTwo(n)
	fmt.Printf("%v", result)
}

func PowersOfTwo(n int) []uint64 {
	// Инициализация слайса для хранения результатов
	var result []uint64

	// Цикл для вычисления степеней 2
	for pow := 0; pow <= n; pow++ {
		// Используем сдвиг, чтобы вычислить 2^pow
		result = append(result, 1<<pow)
	}

	return result
}
