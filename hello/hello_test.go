package main

import "testing"

func TestOla(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, obtido, esperado string) {
		t.Helper()
		if obtido != esperado {
			t.Errorf("obtido '%s', esperado '%s'", obtido, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		esperado := "Olá, Chris!"
		obtido := Ola("Chris", "")

		verificaMensagemCorreta(t, obtido, esperado)
	})

	t.Run("diz 'Olá, Mundo!' quando uma string vazia for passada", func(t *testing.T) {
		esperado := "Olá, Mundo!"
		obtido := Ola("", "")

		verificaMensagemCorreta(t, obtido, esperado)
	})

	t.Run("em espanhol", func(t *testing.T) {
		esperado := "Hola, Elodie!"
		obtido := Ola("Elodie", "espanhol")

		verificaMensagemCorreta(t, obtido, esperado)
	})

	t.Run("em frances", func(t *testing.T) {
		esperado := "Bonjour, Elodie!"
		obtido := Ola("Elodie", "frances")

		verificaMensagemCorreta(t, obtido, esperado)
	})

	t.Run("em alemão", func(t *testing.T) {
		esperado := "Hallo, Elodie!"
		obtido := Ola("Elodie", "alemao")

		verificaMensagemCorreta(t, obtido, esperado)
	})
}
