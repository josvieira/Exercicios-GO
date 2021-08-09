package main

import "fmt"

//funções podem retornar funções
//go possui o conceio de closure, a função respeita o contexto léxico dela
func closure() func() {
	x := 10
	var imprimir = func() {
		fmt.Println(x)
	}
	return imprimir
}

func main() {
	x := 20
	fmt.Println(x)
	funcao := closure()
	funcao() //imprimiu 10 que foi declarado dentro da função closure
}
