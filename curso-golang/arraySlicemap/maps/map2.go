package main

import "fmt"

func main() {
	funcEscolar := map[string]float64{
		"joão":     15465,
		"gabriela": 1545545, //tem que ter a vírgula na última linha ou a chave
	}
	fmt.Println(funcEscolar)

	funcEscolar["Rafael"] = 1586565
	delete(funcEscolar, "Maria") //não dará erro, tenta deletar algo que não existe

}
