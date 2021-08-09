package main

import (
	"fmt"
	"time"
)

func fale(pessoa, frase string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s %s\n", pessoa, frase)
	}
}

//para chamar a funcção de forma concorrente basta colocar a palavra reservada go antes da chamada da funcção
//porém a linha de execução concorrente irá encerrar assim que a linha de execução principal encerrar
func main() {
	go fale("João", "Olá", 10)
	fale("maria", "Oi", 3)

	//imprime algo parecido João Olá
	//João Olá
	// maria Oi
	// João Olá
	// maria Oi
	// maria Oi
}
