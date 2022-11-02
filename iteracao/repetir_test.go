package iteracao

import "testing"

func TestRepetir(t *testing.T) {
	verificaResultado := func(t *testing.T, obtido, esperado string) {
		t.Helper()
		if obtido != esperado {
			t.Errorf("obtido '%s', esperado '%s'", obtido, esperado)
		}
	}

	t.Run("repete 5 vezes", func(t *testing.T) {
		obtido := Repetir("a", 5)
		esperado := "aaaaa"

		verificaResultado(t, obtido, esperado)
	})

	t.Run("repete 10 vezes", func(t *testing.T) {
		obtido := Repetir("a", 10)
		esperado := "aaaaaaaaaa"

		verificaResultado(t, obtido, esperado)
	})

	t.Run("repete 0 vezes", func(t *testing.T) {
		obtido := Repetir("a", 0)
		esperado := ""

		verificaResultado(t, obtido, esperado)
	})

	t.Run("repete -1 vezes", func(t *testing.T) {
		obtido := Repetir("a", -1)
		esperado := ""

		verificaResultado(t, obtido, esperado)
	})

}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 5)
	}
}
