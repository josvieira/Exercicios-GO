package main

import "fmt"

type carro struct {
	nome       string
	velocidade string
}

type ferrari struct {
	carro       //fazendo dessa forma gera campos anônimos dentro da struct ferrari, composição
	turboligado bool
}

func main() {
	f := ferrari{}
	f.turboligado = false
	//a struct ferrari não possui nome e velocidade, mas carro sim, por conta da composição ferrari também passa a ter esses campos
	f.nome = "fusca"
	f.velocidade = "40"
	fmt.Println(f) //imprimi: {{fusca 40} false}

}
