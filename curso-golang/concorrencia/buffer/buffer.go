package main

import "fmt"

func rotina(c chan int) {
	c <- 1
	c <- 2
	c <- 3
	// essa linha será executada após a leitura do valor 1 do buffer e após a inserção do valor 3
	//caso não haja espaço para a inserção do valor trẽs esse print não será executado
	fmt.Println("Executou")
	c <- 4
	fmt.Println("Executou")
}

func main() {
	c := make(chan int, 2)

	go rotina(c)
	fmt.Println(<-c)
}
