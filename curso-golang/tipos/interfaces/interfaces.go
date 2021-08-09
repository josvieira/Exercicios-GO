package main

import "fmt"

type imprimivel interface {
	toString() string
}

type produto struct {
	nome  string
	preco float64
}

type pessoa struct {
	nome      string
	sobrenome string
}

//interfaces são implementadas implicitamente, não preciso colocar implements "nomeDaInterface"
//Basta criar os métodos definidos na interface e associa-los à struct por meio do recive
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	//polimorfiso? coisa agora é produto ao invés de pessoa
	coisa = produto{"calça", 79.30}
	fmt.Println(coisa.toString())
}
