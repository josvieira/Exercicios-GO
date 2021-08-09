package main

import (
	"fmt"
	"time"
)

//channel é a forma de comunicação entre as goroutines
//channel é um tipo nativo da linguagem chan
func doisTresQuatrovezes(base int, c chan int) {
	time.Sleep(time.Second)
	//enquanto o valor que está no canal não é lido o programa não passa para a prima instrução
	//operação bloqueante, pois o canal não tem buffer
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)

	go doisTresQuatrovezes(2, c)

	a := <-c
	fmt.Println(a)
	b := <-c

	fmt.Println(b)
	fmt.Println(<-c)

	//se tentasse fazer a leitura do canal mais uma vez causaria um deadlock
	fmt.Println(<-c)

}
