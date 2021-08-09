package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	//passando um slice para a funçao
	aprovados := []string{"Maria", "João", "Lucas", "pedro"}
	//quando passa um slice para uma funçao variática tem que colocar o spred
	//também não poderia passar um array para uma função variática
	imprimirAprovados(aprovados...)
}
