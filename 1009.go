package main

import (
		"fmt"
)


func main() {
	var nome string
	var salario, vendas float64
	fmt.Scan(&nome, &salario, &vendas)

	comissao := vendas * 0.15

	total := salario + comissao

	fmt.Printf("TOTAL = R$ %.2f\n", total)
}