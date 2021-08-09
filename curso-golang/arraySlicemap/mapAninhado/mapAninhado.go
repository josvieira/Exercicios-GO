package main

import "fmt"

func main() {
	funcionarios := map[string]map[string]float64{
		"g": {
			"Gabriela": 4456565,
			"Gil":      78222,
		},
		"h": {
			"helio": 54545,
		},
		"A": {
			"Ana": 79466,
		},
	}
	for letra, funci := range funcionarios {
		fmt.Println(letra)
		for nome, salario := range funci {
			fmt.Println(nome, salario)
		}
	}
}
