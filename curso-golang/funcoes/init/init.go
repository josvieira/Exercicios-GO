package main

import "fmt"

//por convenção o Go vai executar primeiro a funçaõ init
func init() {
	fmt.Println("Inicializando,,,")
}

func main() {
	fmt.Println("Main...")
}

//imprime primeiro o Inicializando
