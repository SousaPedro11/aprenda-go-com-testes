package arrays

import "testing"

func TestSoma(t *testing.T) {

	verificarResultado := func(t *testing.T, resultado, esperado int) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado %d, esperado %d", resultado, esperado)
		}
	}

	t.Run("coleção de 5 números", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}

		resultado := Soma(numeros)
		esperado := 15

		verificarResultado(t, resultado, esperado)
	})

	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3}

		resultado := Soma(numeros)
		esperado := 6

		verificarResultado(t, resultado, esperado)
	})
}
