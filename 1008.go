package main

import (
		"fmt"
)


func main() {
	var number, hours int
	var value float64
	fmt.Scan(&number, &hours, &value)

	salary := float64(hours) * value

	fmt.Printf("NUMBER = %v\n", number)
	fmt.Printf("SALARY = U$ %.2f\n", salary)
}