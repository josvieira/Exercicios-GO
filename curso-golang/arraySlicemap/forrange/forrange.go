package main

import "fmt"

func main() {
	//não precisa clocar a quantidade de posições caso já seja feita a
	//inicialização do array na declaração, coloca os ... e o compilador "conta"
	numeros := [...]int{1, 2, 3, 4, 5, 6}

	//retorna dois valores o indice i do array e o valor contido naquele indice v
	//podemos ignorar um dos valores colocando o _ no lugar da variável
	//equivalente ao foreach de outras linguagens
	for i, v := range numeros {
		fmt.Printf("Indice:%d = valor:%d\n", i, v)

	}
}
