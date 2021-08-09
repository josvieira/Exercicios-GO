package main

import "fmt"

//go não possui operador ternário =(

func obeterResultado(nota float64) string {
	//return nota >= 6 ? "Aprovada" : "Reprovada"
	if nota >= 6 {
		return "Aprovado"
	} //por uma questão de stilo go recomenda não colocar else aqui
	return "Reprovado"
}

func main() {
	fmt.Println(obeterResultado(6.5))
}
