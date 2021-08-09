package main

import (
	"fmt"
	"time"
)

func isPrimo(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func primos(n int, c chan int) {
	inicio := 2
	for i := n; i < n; i++ {
		for primo := inicio; ; {
			if isPrimo(primo) {
				c <- primo
				fmt.Print(<-c)
				inicio = primo + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	close(c) //após adiconar dados num canal é preciso fecha-lo para não causar dedloacks, a menos que ele seja usado novamente
}

func main() {
	c := make(chan int, 30)

	go primos(cap(c), c)

	for primo := range c {
		fmt.Printf("%d", primo)
	}

	fmt.Println("FIM")
}
