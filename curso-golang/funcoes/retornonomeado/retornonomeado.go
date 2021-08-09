package main

import "fmt"

func trocar(p1, p2 int) (sucesso int, erro int) {
	sucesso = p2
	erro = p1
	return
	/*retorno limpo, uma vez que nomeio os retornos eles serão retornados
	no retorno, se não inicializar as variáveis no escopo da função o valor retornado será o valor de padrão para nulo dependendod do tipo
	*/
}

func main() {
	fmt.Println(trocar(10, 0))

}
