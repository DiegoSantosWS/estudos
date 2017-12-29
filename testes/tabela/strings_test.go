package string

import "testing"
import "strings"

const smgIndex = "%s (parte: %s) - indices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	t.Parallel()
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Cod3r é show", "Cod3r", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"leonardo", "o", 2},
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)
		if atual != teste.esperado {
			t.Errorf(smgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}
