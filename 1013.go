package main

import (
	"fmt"
	"math"
)

func maior(a int, b int) int {
	x := a - b
	maior := (a + b + int(math.Abs(float64(x)))) / 2

	return maior
}

func main() {
	var A, B, C int

	fmt.Scan(&A, &B, &C)
	res := maior(A, B)

	resFinal := maior(res, C)

	fmt.Printf("%v eh o maior\n", resFinal)
}