package main

import "fmt"

//retorna mais de uma variável
func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprartv50 := trab1 && trab2
	comprartv32 := trab1 != trab2 //ou exclusivo
	comprarsorvete := trab1 || trab2

	return comprartv50, comprartv32, comprarsorvete
}

func main() {
	fmt.Println(compras(true, false))
}
