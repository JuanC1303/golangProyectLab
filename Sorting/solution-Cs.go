package main

import "fmt"

func main() {
	fmt.Println(countingSort([]int{10, 2, 3, 7, 50, 34, 12, 33, 35, 99, 65, 12, 67, 12, 77, 32, 78, 89, 54, 34, 55, 33, 66}))
}

func countingSort(slice []int) []int {
	if len(slice) == 0 {
		return slice
	}
	tocount := cZeros(fMin(slice),
		fMax(slice))

	for _, i := range slice {
		tocount[i-fMin(slice)] += 1
	}
	for i := 1; i < len(tocount); i++ {
		tocount[i] += tocount[i-1]
	}

	res := make([]int, len(slice))

	for i := 0; i < len(slice); i++ {
		j := slice[i]
		k := tocount[j-fMin(slice)] - 1

		res[k] = j
		tocount[j-fMin(slice)] = tocount[j-fMin(slice)] - 1
	}
	return res
}

func fMax(slice []int) int {

	max := slice[0]

	for _, i := range slice {
		if max < i {
			max = i
		}
	}
	return max
}

func fMin(slice []int) int {

	min := slice[0]

	for _, i := range slice {
		if i < min {
			min = i
		}
	}
	return min
}

func cZeros(min, max int) []int {
	zeros := make([]int, max-min+1)

	for i := range zeros {
		zeros[i] = 0
	}
	return zeros
}
