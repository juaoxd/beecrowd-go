package main

import (
		"fmt"
)


func main() {
	var A, B int
	fmt.Scan(&A, &B)

	prod := A * B

	fmt.Printf("PROD = %v\n", prod)
}