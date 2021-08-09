package main

import "fmt"

func main() {
	i := 1

	//go não possui aritmetica de ponteiros
	//declarando um ponterio
	var p *int = nil

	//fazendo o ponteiro p apontar para o endereço de memória de i
	p = &i

	fmt.Println(p) //o conteudo de p é o endereço de memória de i
	fmt.Println(&i)
	fmt.Println(&p)      //aqui imprime o endereço de memória do pŕoprio ponteiro, onde ele está alocado
	fmt.Println(*p)      //aqui imprime o conteúdo do endereço de memória para o qual ele aponta no caso a variável i
	fmt.Println(p == &i) //true
	fmt.Println(*p == i) //true
}
