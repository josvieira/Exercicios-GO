package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	//posso usar interface como um tipo genério, como se fosse o Object do Java? talvez
	//coisa e coisa2 vao poder receber qualquer tipo de dado
	//como usar isso? qual finalidade ainda não sei
	var coisa interface{}
	type dinamico interface{}
	var coisa2 dinamico = "Opa"

	fmt.Println(coisa) //imprimi nil
	coisa = 2.5
	fmt.Println(coisa) //imprime 2.5

	fmt.Println(coisa2) //imprime opa
	coisa2 = true
	fmt.Println(coisa2) //imprime true
}
