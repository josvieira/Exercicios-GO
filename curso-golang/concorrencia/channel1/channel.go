package main

import "fmt"

func main() {

	//canal criado com buffer de 1
	ch := make(chan int, 1)

	ch <- 1 //passando/escrevendo o valor 1 para o canal
	<-ch    //lemdo o valor de um canal

	ch <- 2

	fmt.Println(<-ch)
}
