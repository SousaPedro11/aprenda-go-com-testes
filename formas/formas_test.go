package formas

import "testing"

func TestPerimetro(t *testing.T) {

	testesPerimetro := []struct {
		nome     string
		forma    Forma
		esperado float64
	}{
		{nome: "Retângulo", forma: Retangulo{Altura: 12, Base: 6}, esperado: 36.0},
		{nome: "Círculo", forma: Circulo{Raio: 10}, esperado: 62.83185307179586},
		{nome: "Triângulo", forma: Triangulo{Base: 3, Altura: 4}, esperado: 12.0},
	}

	verificarResultado := func(t *testing.T, resultado, esperado float64) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%.2f', esperado '%.2f'", resultado, esperado)
		}
	}

	for _, tt := range testesPerimetro {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := tt.forma.Perimetro()
			verificarResultado(t, resultado, tt.esperado)
		})
	}
}

func TestArea(t *testing.T) {

	testesArea := []struct {
		nome     string
		forma    Forma
		esperado float64
	}{
		{nome: "Retângulo", forma: Retangulo{Altura: 12, Base: 6}, esperado: 72.0},
		{nome: "Círculo", forma: Circulo{Raio: 10}, esperado: 314.1592653589793},
		{nome: "Triângulo", forma: Triangulo{Base: 12, Altura: 6}, esperado: 36.0},
	}

	verificarResultado := func(t *testing.T, resultado, esperado float64) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado %.2f esperado %.2f", resultado, esperado)
		}
	}

	for _, tt := range testesArea {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := tt.forma.Area()
			verificarResultado(t, resultado, tt.esperado)
		})
	}
}
