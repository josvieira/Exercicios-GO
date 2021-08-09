package strings

import (
	"strings"
	"testing"
)

const erroPadrao = "O valor esperado é %v, mas vaio %v"

func TestIndex(t *testing.T) {
	t.Parallel() //indica que os testes podem rodar em paralelo
	dataSet := []struct {
		texto    string
		parte    string
		esperado int
	}{
		//data set criado para testar várias possibilidades da função
		{"Coder Br Cursos", "Coder", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"leo", "o", 2},
	}

	for _, teste := range dataSet {
		t.Logf("massa")
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(erroPadrao, teste.esperado, atual)
		}
	}
}
