package main

import (
	"fmt"
	"strconv"
)

func main() {
	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//Aqui ele não converte 97 numa string mas pega o código da tabela ascii a
	fmt.Println("teste" + string(97))

	//para transformar int em string o contrario o método é Atoi
	fmt.Println("Teste " + strconv.Itoa(97))

	//A função Atoi retorna dois valores, o segundo é um valor de erro, podemos usar ou ignorar com _
	num, _ := strconv.Atoi("2")
	fmt.Println(num)

	//aceita apenas string para transformar em boolean
	bo, _ := strconv.ParseBool("1")
	if bo {
		fmt.Println(bo)
	}
}
