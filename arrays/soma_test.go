package arrays

import (
	"reflect"
	"testing"
)

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

func TestSomaTudo(t *testing.T) {
	resultado := SomaTudo([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}

	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("resultado %v, esperado %v", resultado, esperado)
	}
}

func TestSomaTodoOResto(t *testing.T) {
	verificarSomas := func(t *testing.T, resultado, esperado []int) {
		t.Helper()
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado %v, esperado %v", resultado, esperado)
		}
	}

	t.Run("faz a soma dos elementos", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
		esperado := []int{2, 9}
		verificarSomas(t, resultado, esperado)
	})

	t.Run("soma de um slice vazio", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{}, []int{3, 4, 5})
		esperado := []int{0, 9}
		verificarSomas(t, resultado, esperado)
	})
}
