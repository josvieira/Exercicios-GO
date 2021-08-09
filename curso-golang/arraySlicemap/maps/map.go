package main

import "fmt"

func main() {
	//var aprovados map[int]string //gera um map nulo devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[04545] = "Maria"
	aprovados[04545] = "Marta" //não aceita chave duplicada, substitui pelo ultimo valor
	aprovados[24555] = "João"
	aprovados[5454554] = "Luiza"
	aprovados[54765465] = "Lucas"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s - CPF: %d\n", nome, cpf)
	}

}
