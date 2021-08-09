package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	switch { //um switch true e vai executar na primeira condição verdadeira
	case t.Hour() > 10:
		fmt.Println("Tarde")
	case t.Hour() < 10:
		fmt.Println("Cedo")

	}
}
