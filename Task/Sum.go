package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// "os" se utiliza para escribir los datos desde la linea de comando que corre el programa.
	// length := 0
	// fmt.Println("Enter the number of inputs")
	// fmt.Scanln(&length)
	// fmt.Println("Enter the inputs")
	// x := make([]int, length)
	// for i := 0; i < length; i++ {
	// 	fmt.Scanln(&x[i])
	// }

	str := os.Args[1:]
	x := make([]int, len(str))

	for i, err := range str {
		x[i], _ = strconv.Atoi(err)
	}

	fmt.Println(x)

	var p int

	p = negative(x)
	fmt.Println("The final result is:", sum1(p, x))
}

func negative(x []int) int {

	for i, v := range x {
		if v < 0 {
			return (i - 1)
		}
	}
	return 0
}

func sum1(p int, x []int) int {
	sum := 0

	if len(x) == 0 {
		return 0
	}

	for i := 0; i <= p; i++ {
		if i > 0 && i < p {
			sum += x[i] * 2
		} else {
			sum += x[i]
		}
	}
	return sum
}
