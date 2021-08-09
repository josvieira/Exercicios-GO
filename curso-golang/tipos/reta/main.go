package main

import "fmt"

//como temos mais de um arquivo para executar é go run *.go
func main() {
	p1 := Ponto{2.0, 4.0}
	p2 := Ponto{2.0, 8.0}

	//cateto está minusculo é privado mas visível no pacote
	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1, p2))
}
