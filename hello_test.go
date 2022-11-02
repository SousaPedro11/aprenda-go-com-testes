package main

import "testing"

func TestOla(t *testing.T) {
	obtido := Ola("Pedro")
	esperado := "Olá, Pedro!"

	if obtido != esperado {
		t.Errorf("obtido '%s' esperado '%s'", obtido, esperado)
	}
}
