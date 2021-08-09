package main

import (
	"fmt"
	"math"
)

func main() {
	a := 4
	b := 2
	//operações binários iguais as outras linguagens

	//operadores bitwise, os resultados são valores numericos e não valores lógicos
	fmt.Println("E =>", a&b)   //0?
	fmt.Println("Ou =>", a|b)  //6?
	fmt.Println("Xor =>", a^b) //6?

	c := 2.0
	d := 4.0
	//operações com math
	fmt.Println("maior que =>", math.Max(float64(a), float64(b)))
	fmt.Println("potência =>", math.Pow(c, d))
}
