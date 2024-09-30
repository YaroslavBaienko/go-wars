package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Compare a nd b")
	a := []int{2, 1, 9}
	b := []int{81, 4, 1}
	result := Comp(a, b)
	fmt.Printf("%v result", result)
}

func Comp(array1 []int, array2 []int) bool {
	// Check if either array is nil
	if array1 == nil || array2 == nil {
		return false
	}

	// Check if arrays are of different length
	if len(array1) != len(array2) {
		return false
	}

	// Create copies of the arrays to avoid modifying the originals
	cpyArray1 := make([]int, len(array1))
	cpyArray2 := make([]int, len(array2))
	copy(cpyArray1, array1)
	copy(cpyArray2, array2)

	// Sort both arrays for proper comparison
	sort.Ints(cpyArray1)
	sort.Ints(cpyArray2)

	// Square elements in cpyArray1 and compare with elements in cpyArray2
	for i, v := range cpyArray1 {
		if v*v != cpyArray2[i] {
			return false
		}
	}

	return true
}
