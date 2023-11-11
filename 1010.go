package main

import (
	"fmt"
)

func main() {
	var codigo, qtd int
	var valor float64

	fmt.Scan(&codigo, &qtd, &valor)
	total := float64(qtd) * valor

	fmt.Scan(&codigo, &qtd, &valor)
	total += float64(qtd) * valor
	

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}