package main

import (
	"fmt"
)

//go suporta recursividade
//posso utilizar aqui inteiros sem sinal não preciso testar uma das condições do switch
func fatorial(num uint) uint {
	//debug.PrintStack()
	switch {
	case num == 0:
		return 1
	default:
		return num * fatorial(num-1)
	}

}

func main() {
	//fmt.Println(fatorial(-5))
	fmt.Println(fatorial(5))

}
