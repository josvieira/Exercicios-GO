package main

import "fmt"

type item struct {
	produtoId  int
	preco      float64
	quantidade int
}

type pedido struct {
	pedidoId int
	intens   []item
}

func (p pedido) calcularValor() float64 {
	total := 0.0
	for _, item := range p.intens {
		total += item.preco * float64(item.quantidade)
	}
	return total
}

func main() {
	pedido := pedido{
		pedidoId: 1,
		intens: []item{
			item{1, 5.8, 5},
			item{2, 6.8, 6},
		},
	}
	fmt.Println(pedido.calcularValor())
	fmt.Println(pedido)
}
