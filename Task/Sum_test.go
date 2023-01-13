package main

import "testing"

func testSum(t *testing.T) {

	if sum1([]int{1, 2, 3, 4, 5, -6}) != 24 {
		t.Fatal("The final result is: 24")
	}

}
