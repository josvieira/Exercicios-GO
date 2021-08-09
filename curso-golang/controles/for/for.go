package main

import (
	"fmt"
	"time"
)

func main() {
	//go não tem while apenas for

	i := 1
	for i != 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i++ {
		fmt.Println(i)
	}

	//laço infinito
	for {
		fmt.Println("Para sempre")
		time.Sleep(time.Second)
	}
}
