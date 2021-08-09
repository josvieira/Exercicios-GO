package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNome() string {
	return p.nome + " " + p.sobrenome
}

//Toda passagem de pearametro para função é por valor, se quero alterar um
//campo da struct preciso fazer a passagem do parãmetro por referencia recebendo um ponteiro
func (p *pessoa) setNome(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{"João", "Silva"}
	fmt.Println(p1.getNome())

	//Aqui na hora de chamar a função não precisa passar a referencia de
	//memória do ponteiro &p1, o próprio Go faz isso
	p1.setNome("Maria Antônia")
	fmt.Println(p1.getNome())
}
