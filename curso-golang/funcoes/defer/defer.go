package main

import "fmt"

func obetervalor(num int) int {
	//defer faz com que essa instrução de código seja executada antes do return da funçaõ
	defer fmt.Println("Fim")
	if num > 500 {
		fmt.Println("Valor alto")
		return num
	} else {
		fmt.Println("Valor Baixo")
		return num
	}

}

func main() {
	fmt.Println(obetervalor(500))
}
