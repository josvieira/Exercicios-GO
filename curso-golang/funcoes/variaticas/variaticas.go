package main

import "fmt"

//funções variáticas são as que recebem quantidades de parãmetros variáveis
//algunas linguagens usam o var args, o operador spreed
func media(numeros ...float64) float64 {
	total := 0.0
	for _, valor := range numeros {
		total += valor
	}
	return total / float64(len(numeros))
}

func main() {
	//posso chamar a funçaõ média com quanquer quantidade de parâmetros
	fmt.Printf("Media: %.3f", media(5.2, 5.5, 8.6, 9.1))
}
