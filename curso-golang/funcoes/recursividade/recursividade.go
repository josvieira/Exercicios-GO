package main

import (
	"fmt"
	"runtime/debug"
)

//go suporta recursividade
func fatorial(num int) (int, error) {
	debug.PrintStack()
	switch {
	case num < 0:
		return -1, fmt.Errorf("Número inválido: %d", num)
	case num == 0:
		return 1, nil
	default:
		fatorialanteriot, _ := fatorial(num - 1)
		return num * fatorialanteriot, nil
	}

}

func main() {
	fmt.Println(fatorial(5))

}
