package main

import (
	"fmt"
	f "fmt"
)

func main() {
	const PI float64 = 3.1415
	//se um valor for atribuido a uma variável não precisa declara o tipo, o go faz a inferencia
	var raio = 3.1
	f.Print(raio)
	fmt.Printf("\n%f ", raio) //semelhante ao C %d, %f %.2f %t %s %v
	fmt.Println("\no raio é ", raio)

	//se usar com igual que dizer que a variavel raio já existe
	raio = 2
	raio = "string" // não posso fazer isso pois a tipagem é estática pela primeira declaração da variável
	//forma reduzida de declara variável, com o : significa que ela ainda não existe
	area := 5 //sinalizando erro de compilação pois não estamos usando a variável

	const (
		a = 1
		b = 2
	)

	var (
		c = 2
		d = 3
	)
}
