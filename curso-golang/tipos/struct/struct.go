package main

import "fmt"

//a forma de declarar uma struct
type produto struct {
	nome     string
	preco    float64
	desconto float64
}

//Método: conceito de função receiver(receptor), uma função que é associada a um tipo
//essa função agora está associada à struct produto posso fazer produto.recoComDesconto
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	//duas formas diferentes que criar e inicializar uma struct
	var produto1 produto
	produto1 = produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	produto2 := produto{"Notebook", 1.68800, 0.15}

	fmt.Println(produto1.nome, produto1.precoComDesconto())
	fmt.Println(produto2.nome, produto2.precoComDesconto())
}
