package main

import "testing"

func TestSum(t *testing.T) {

	if sum1(4, []int{2, 1, 3, 4, 5, -5}) != 23 {
		t.Fatal("The final result is: 23")
	}

	if sum1(0, []int{}) != 0 {
		t.Fatal("The final result is: 0")
	}

	if sum1(0, []int{2, -5}) != 2 {
		t.Fatal("The final result is: 2")
	}

	// if sum1(0, []int{-2, -1, -3, -4, -5, -5}) != 0 {
	// 	t.Fatal("The final result is: 0")
	// }

	// if sum1(4, []int{2, 1, 3, 4, 5, 5}) != 0 {
	// 	t.Fatal("The final result is: 0")
	// }

}
