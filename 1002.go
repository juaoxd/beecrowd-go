package main

import (
		"fmt"
		"math"
)


func main() {
	const pi = 3.14159
	var raio float64

	fmt.Scan(&raio)

	area := math.Pow(raio, 2) * pi

	fmt.Printf("A=%.4f\n", area)

}