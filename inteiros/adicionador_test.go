package inteiros

import "testing"

func TestAdicionador(t *testing.T) {

	verificaResultado := func(t *testing.T, obtido, esperado int) {
		t.Helper()
		if obtido != esperado {
			t.Errorf("obtido '%d', esperado '%d'", obtido, esperado)
		}
	}

	t.Run("adiciona 0 + 0 e obtém 0", func(t *testing.T) {
		obtido := Adiciona(0, 0)
		esperado := 0

		verificaResultado(t, obtido, esperado)
	})

	t.Run("adiciona 0 + 2 e obtém 2", func(t *testing.T) {
		obtido := Adiciona(0, 2)
		esperado := 2

		verificaResultado(t, obtido, esperado)
	})

	t.Run("adiciona 2 + 2 e obtém 4", func(t *testing.T) {
		obtido := Adiciona(2, 2)
		esperado := 4

		verificaResultado(t, obtido, esperado)
	})

	t.Run("adiciona -1 + 2 e obtém 1", func(t *testing.T) {
		obtido := Adiciona(-1, 2)
		esperado := 1

		verificaResultado(t, obtido, esperado)
	})

	t.Run("adiciona 1 + -2 e obtém -1", func(t *testing.T) {
		obtido := Adiciona(1, -2)
		esperado := -1

		verificaResultado(t, obtido, esperado)
	})

	t.Run("adiciona -1 + -2 e obtém -3", func(t *testing.T) {
		obtido := Adiciona(-1, -2)
		esperado := -3

		verificaResultado(t, obtido, esperado)
	})

}
