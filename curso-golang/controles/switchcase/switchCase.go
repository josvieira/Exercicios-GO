package main

func notaParaConceito(n float64) string {
	var nota = int(n)

	switch nota { //aqui o switch não precisa de break para sair de um case
	case 10:
		fallthrough //ele precisa dessa palavra caso seja para seja para o próximo bloco
	case 9:
		return "A"
	case 8, 7: //para o caso de mais de uma opção tiver o mesmo resutado
		return "B"
	default:
		return "Nota Inválida"
	}
}
