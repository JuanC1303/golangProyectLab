package main

import (
	"fmt"
)

func main() {
	fmt.Println(insertionSort([]int{2, 4, 7, 1, 5}))
}

func insertionSort(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i; j > 0 && slice[j-1] > slice[j]; j-- {
			slice[j], slice[j-1] = slice[j-1], slice[j]
		}
	}
	return slice
}
