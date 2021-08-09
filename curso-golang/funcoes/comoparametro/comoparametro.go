package main

import "fmt"

func multiplicacao(num1, num2 int) int {
	return num1 * num2
}

//uma funçaõ pode ser passada como parâmetro
//pesquisar exemplos para treinar isso
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	fmt.Println(exec(multiplicacao, 2, 3))
}
