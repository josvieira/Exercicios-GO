package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(2 * time.Second)
	c <- 1 //operação bloqueante
}

func main() {
	c := make(chan int)

	go rotina(c)

	fmt.Println(<-c) //operação bloqueante fica aqui esperando o dado ser lido
	fmt.Println("Foi lido")
}
